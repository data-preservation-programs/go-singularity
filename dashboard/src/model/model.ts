// Code generated by model2ts. DO NOT EDIT.
export type WorkType = 'scan' | 'deal_making' | 'deal_tracking' | 'packing'
export type WorkState = '' | 'ready' | 'processing' | 'complete' | 'error'
export type DealState = 'proposed' | 'published' | 'active' | 'expired' | 'proposal_expired' | 'rejected' | 'slashed' | 'error'
export type ScheduleState = 'active' | 'paused' | 'error' | 'completed'
export interface Global {
	key: string
	value: string
}
export interface Worker {
	id: string
	workType: WorkType
	workingOn: string
	lastHeartbeat: Date
	hostname: string
}
export interface Dataset {
	id: number
	name: string
	createdAt: Date
	updatedAt: Date
	maxSize: number
	pieceSize: number
	outputDirs: string[]
	encryptionRecipients: string[]
	encryptionScript: string
	wallets: Wallet[] | undefined
}
export interface Source {
	id: number
	createdAt: Date
	updatedAt: Date
	datasetId: number
	dataset: Dataset | undefined
	type: string
	path: string
	metadata: {[key: string]: string}
	scanIntervalSeconds: number
	scanningState: WorkState
	scanningWorkerId: string | undefined
	scanningWorker: Worker | undefined
	lastScannedTimestamp: number
	lastScannedPath: string
	errorMessage: string
	deleteAfterExport: boolean
	dagGenState: WorkState
	dagGenWorkerId: string | undefined
	dagGenWorker: Worker | undefined
	dagGenErrorMessage: string
}
export interface Chunk {
	id: number
	createdAt: Date
	sourceId: number
	source: Source | undefined
	packingState: WorkState
	packingWorkerId: string | undefined
	packingWorker: Worker | undefined
	errorMessage: string
	itemParts: ItemPart[] | undefined
	cars: Car[] | undefined
}
export interface Item {
	id: number
	createdAt: Date
	sourceId: number
	source: Source | undefined
	path: string
	hash: string
	size: number
	lastModified: number
	cid: string
	directoryId: number
	directory: Directory | undefined
	itemParts: ItemPart[] | undefined
}
export interface Directory {
	id: number
	updatedAt: Date
	cid: string
	sourceId: number
	source: Source | undefined
	name: string
	parentId: number
	parent: Directory | undefined
	exported: boolean
}
export interface Car {
	id: number
	createdAt: Date
	pieceCid: string
	pieceSize: number
	rootCid: string
	fileSize: number
	filePath: string
	datasetId: number
	dataset: Dataset | undefined
	sourceId: number
	source: Source | undefined
	chunkId: number
	chunk: Chunk | undefined
	header: string
}
export interface CarBlock {
	id: number
	carId: number
	car: Car | undefined
	cid: string
	carOffset: number
	carBlockLength: number
	varint: string
	rawBlock: string
	itemId: number
	item: Item | undefined
	itemOffset: number
	itemEncrypted: boolean
}
export interface Deal {
	id: number
	createdAt: Date
	updatedAt: Date
	dealId: number
	datasetId: number
	dataset: Dataset | undefined
	state: DealState
	clientId: string
	wallet: Wallet | undefined
	provider: string
	proposalId: string
	label: string
	pieceCid: string
	pieceSize: number
	startEpoch: number
	endEpoch: number
	sectorStartEpoch: number
	price: string
	verified: boolean
	errorMessage: string
	scheduleId: number
	schedule: Schedule | undefined
}
export interface Schedule {
	id: number
	createdAt: Date
	updatedAt: Date
	datasetId: number
	dataset: Dataset | undefined
	urlTemplate: string
	httpHeaders: string[]
	provider: string
	pricePerGbEpoch: number
	pricePerGb: number
	pricePerDeal: number
	totalDealNumber: number
	totalDealSize: number
	verified: boolean
	keepUnsealed: boolean
	announceToIpni: boolean
	startDelay: number
	duration: number
	state: ScheduleState
	scheduleCron: string
	scheduleDealNumber: number
	scheduleDealSize: number
	maxPendingDealNumber: number
	maxPendingDealSize: number
	notes: string
	errorMessage: string
	allowedPieceCids: string[]
}
export interface Wallet {
	id: string
	address: string
	privateKey: string | undefined
	remotePeer: string | undefined
}
export interface WalletAssignment {
	id: number
	walletId: string
	wallet: Wallet | undefined
	datasetId: number
	dataset: Dataset | undefined
}
export interface ItemPart {
	id: number
	itemId: number
	item: Item | undefined
	offset: number
	length: number
	cid: string
	chunkId: number
	chunk: Chunk | undefined
}