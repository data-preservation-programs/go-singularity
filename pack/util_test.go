package pack

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/model"
	"github.com/ipfs/boxo/util"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"github.com/rclone/rclone/backend/s3"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/hash"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateParentNode(t *testing.T) {
	links := []format.Link{
		{
			Name: "",
			Size: 5,
			Cid:  cid.NewCidV1(cid.Raw, util.Hash([]byte("hello"))),
		},
		{
			Name: "",
			Size: 5,
			Cid:  cid.NewCidV1(cid.Raw, util.Hash([]byte("world"))),
		},
	}

	node, size, err := createParentNode(links)
	require.NoError(t, err)
	require.Equal(t, uint64(10), size)
	require.NotNil(t, node)
	size, err = node.Size()
	require.NoError(t, err)
	require.Equal(t, uint64(108), size)
	require.Equal(t, "bafybeiejlvvmfokp5c6q2eqgbfjeaokz3nqho5c7yy3ov527vsatgsqfma", node.String())
}

func TestMin(t *testing.T) {
	require.Equal(t, 1, Min(1, 2))
	require.Equal(t, 1, Min(2, 1))
	require.Equal(t, 1, Min(1, 1))
}

func TestAssembleItemFromLinks(t *testing.T) {
	links := []format.Link{
		{
			Name: "",
			Size: 5,
			Cid:  cid.NewCidV1(cid.Raw, util.Hash([]byte("hello"))),
		},
		{
			Name: "",
			Size: 5,
			Cid:  cid.NewCidV1(cid.Raw, util.Hash([]byte("world"))),
		},
	}

	blocks, node, err := AssembleItemFromLinks(links)
	require.NoError(t, err)
	require.NotNil(t, node)
	size, err := node.Size()
	require.NoError(t, err)
	require.Equal(t, uint64(108), size)
	require.Equal(t, "bafybeiejlvvmfokp5c6q2eqgbfjeaokz3nqho5c7yy3ov527vsatgsqfma", node.String())
	require.Len(t, blocks, 1)
	require.Equal(t, "bafybeiejlvvmfokp5c6q2eqgbfjeaokz3nqho5c7yy3ov527vsatgsqfma", blocks[0].Cid().String())
}

func TestAssembleItemFromLinks_LargeFile(t *testing.T) {
	links := []format.Link{}
	for i := 0; i < 2_000; i++ {
		links = append(links, format.Link{
			Name: "",
			Size: 5,
			Cid:  cid.NewCidV1(cid.Raw, util.Hash([]byte("hello"))),
		})
	}

	blocks, node, err := AssembleItemFromLinks(links)
	require.NoError(t, err)
	require.NotNil(t, node)
	size, err := node.Size()
	require.NoError(t, err)
	require.EqualValues(t, 10103, size)
	require.Equal(t, "bafybeiamlrjlfotypfc5hse7uenmgmav7yq5vcb75yd7bxh2aaavxbi4ou", node.String())
	require.Len(t, blocks, 3)
	require.Equal(t, "bafybeidci5xcdskgvefhv3x6kp6zpyxkbiqshqjbpsdgmk2k3mexzaggwi", blocks[0].Cid().String())
}

func TestGenerateCarHeader(t *testing.T) {
	header, err := GenerateCarHeader(cid.NewCidV1(cid.Raw, util.Hash([]byte("hello"))))
	require.NoError(t, err)
	require.Len(t, header, 59)
}

func TestWriteCarHeader(t *testing.T) {
	buf := new(bytes.Buffer)
	header, err := WriteCarHeader(buf, cid.NewCidV1(cid.Raw, util.Hash([]byte("hello"))))
	require.NoError(t, err)
	require.Len(t, header, 59)
	require.Equal(t, buf.Bytes(), header)
}

func TestWriteCarBlock(t *testing.T) {
	buf := new(bytes.Buffer)
	block := blocks.NewBlock([]byte("hello"))
	n, err := WriteCarBlock(buf, block)
	require.NoError(t, err)
	require.Equal(t, int64(40), n)
	require.Equal(t, 40, buf.Len())
}

type MockReadHandler struct {
	mock.Mock
}

func (m *MockReadHandler) Read(ctx context.Context, path string, offset int64, length int64) (io.ReadCloser, fs.Object, error) {
	args := m.Called(ctx, path, offset, length)
	if args.Get(1) == nil {
		return args.Get(0).(io.ReadCloser), nil, args.Error(2)
	}
	return args.Get(0).(io.ReadCloser), args.Get(1).(fs.Object), args.Error(2)
}

func TestGetBlockStreamFromItem(t *testing.T) {
	ctx := context.Background()
	mockObject := new(MockObject)
	sizeCall := mockObject.On("Size").Return(int64(5))
	mockObject.On("Fs").Return(&s3.Fs{})
	mockObject.On("Hash", mock.Anything, mock.Anything).Return("hash", nil)
	tm := time.Now()
	mockObject.On("ModTime", mock.Anything).Return(tm)
	handler := new(MockReadHandler)
	readCall := handler.On("Read", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(io.NopCloser(bytes.NewReader([]byte("hello"))), mockObject, nil)

	t.Run("size mismatch", func(t *testing.T) {
		item := model.ItemPart{
			Offset: 0,
			Length: 5,
			Item: &model.Item{
				Size:                      4,
				Hash:                      "hash",
				LastModifiedTimestampNano: tm.UnixNano(),
			},
		}
		_, _, err := GetBlockStreamFromItem(ctx, handler, item, nil)
		require.ErrorIs(t, err, ErrItemModified)
	})

	t.Run("success", func(t *testing.T) {
		item := model.ItemPart{
			Offset: 0,
			Length: 5,
			Item: &model.Item{
				Size:                      5,
				Hash:                      "hash",
				LastModifiedTimestampNano: tm.UnixNano(),
			},
		}
		blockResultChan, _, err := GetBlockStreamFromItem(ctx, handler, item, nil)
		require.NoError(t, err)
		blockResults := make([]BlockResult, 0)
		for r := range blockResultChan {
			blockResults = append(blockResults, r)
		}
		require.Len(t, blockResults, 1)
		require.EqualValues(t, 0, blockResults[0].Offset)
		require.Equal(t, []byte("hello"), blockResults[0].Raw)
		require.Equal(t, "bafkreibm6jg3ux5qumhcn2b3flc3tyu6dmlb4xa7u5bf44yegnrjhc4yeq", blockResults[0].CID.String())
	})

	t.Run("success with empty item", func(t *testing.T) {
		sizeCall.Unset()
		mockObject.On("Size").Return(int64(0))
		readCall.Unset()
		handler.On("Read", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(io.NopCloser(bytes.NewReader([]byte(""))), mockObject, nil)
		item := model.ItemPart{
			Offset: 0,
			Length: 0,
			Item: &model.Item{
				Size:                      0,
				Hash:                      "hash",
				LastModifiedTimestampNano: tm.UnixNano(),
			},
		}
		blockResultChan, _, err := GetBlockStreamFromItem(ctx, handler, item, nil)
		require.NoError(t, err)
		blockResults := make([]BlockResult, 0)
		for r := range blockResultChan {
			blockResults = append(blockResults, r)
		}
		require.Len(t, blockResults, 1)
		require.EqualValues(t, 0, blockResults[0].Offset)
		require.Equal(t, []byte(nil), blockResults[0].Raw)
		require.Equal(t, "bafkreihdwdcefgh4dqkjv67uzcmw7ojee6xedzdetojuzjevtenxquvyku", blockResults[0].CID.String())
	})
}

func TestIsSameEntry(t *testing.T) {
	ctx := context.Background()
	mockObject := new(MockObject)
	mockObject.On("Size").Return(int64(5))
	mockObject.On("Fs").Return(&s3.Fs{})
	mockObject.On("Hash", mock.Anything, mock.Anything).Return("hash", nil)
	tm := time.Now()
	mockObject.On("ModTime", mock.Anything).Return(tm)
	t.Run("size mismatch", func(t *testing.T) {
		same, detail := IsSameEntry(ctx, model.Item{
			Size: 4,
		}, mockObject)
		require.False(t, same)
		require.Contains(t, detail, "size mismatch")
	})
	t.Run("hash mismatch", func(t *testing.T) {
		same, detail := IsSameEntry(ctx, model.Item{
			Size: 5,
			Hash: "hash2",
		}, mockObject)
		require.False(t, same)
		require.Contains(t, detail, "hash mismatch")
	})
	t.Run("last modified mismatch", func(t *testing.T) {
		same, detail := IsSameEntry(ctx, model.Item{
			Size:                      5,
			Hash:                      "hash",
			LastModifiedTimestampNano: 100,
		}, mockObject)
		require.False(t, same)
		require.Contains(t, detail, "last modified mismatch")
	})
	t.Run("all match", func(t *testing.T) {
		same, _ := IsSameEntry(ctx, model.Item{
			Size:                      5,
			Hash:                      "hash",
			LastModifiedTimestampNano: tm.UnixNano(),
		}, mockObject)
		require.True(t, same)
	})
	t.Run("all match, ignoring empty hash", func(t *testing.T) {
		same, _ := IsSameEntry(ctx, model.Item{
			Size:                      5,
			LastModifiedTimestampNano: tm.UnixNano(),
		}, mockObject)
		require.True(t, same)
	})
}

type MockObject struct {
	mock.Mock
}

func (m *MockObject) Remote() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockObject) ModTime(ctx context.Context) time.Time {
	args := m.Called(ctx)
	return args.Get(0).(time.Time)
}

func (m *MockObject) Size() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

func (m *MockObject) Fs() fs.Info {
	args := m.Called()
	return args.Get(0).(fs.Info)
}

func (m *MockObject) Hash(ctx context.Context, ty hash.Type) (string, error) {
	args := m.Called(ctx, ty)
	return args.String(0), args.Error(1)
}

func (m *MockObject) Storable() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockObject) SetModTime(ctx context.Context, t time.Time) error {
	args := m.Called(ctx, t)
	return args.Error(0)
}

func (m *MockObject) Open(ctx context.Context, options ...fs.OpenOption) (io.ReadCloser, error) {
	args := m.Called(ctx, options)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}

func (m *MockObject) Update(ctx context.Context, in io.Reader, src fs.ObjectInfo, options ...fs.OpenOption) error {
	args := m.Called(ctx, in, src, options)
	return args.Error(0)
}

func (m *MockObject) Remove(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}
