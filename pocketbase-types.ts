/**
* This file was @generated using pocketbase-typegen
*/

export enum Collections {
	BookmarkTags = "bookmark_tags",
	Bookmarks = "bookmarks",
	Collections = "collections",
	Groups = "groups",
	Tags = "tags",
	Users = "users",
}

// Alias types for improved usability
export type IsoDateString = string
export type RecordIdString = string

// System fields
export type BaseSystemFields = {
	id: RecordIdString
	created: IsoDateString
	updated: IsoDateString
	collectionId: string
	collectionName: Collections
	expand?: { [key: string]: any }
}

export type AuthSystemFields = {
	email: string
	emailVisibility: boolean
	username: string
	verified: boolean
} & BaseSystemFields

// Record types for each collection

export type BookmarkTagsRecord = {
	bookmark: RecordIdString
	tag: RecordIdString
}

export type BookmarksRecord = {
	user: RecordIdString
	collection: RecordIdString
	url: string
	favourite: boolean
	custom_order?: number
}

export type CollectionsRecord = {
	parent?: RecordIdString
	name: string
	user: RecordIdString
	group?: RecordIdString
	custom_order?: number
}

export type GroupsRecord = {
	user?: RecordIdString
	name?: string
	custom_order?: number
}

export type TagsRecord = {
	tag: string
}

export type UsersRecord = {
	name?: string
	avatar?: string
}

// Response types include system fields and match responses from the PocketBase API
export type BookmarkTagsResponse = BookmarkTagsRecord & BaseSystemFields
export type BookmarksResponse = BookmarksRecord & BaseSystemFields
export type CollectionsResponse = CollectionsRecord & BaseSystemFields
export type GroupsResponse = GroupsRecord & BaseSystemFields
export type TagsResponse = TagsRecord & BaseSystemFields
export type UsersResponse = UsersRecord & AuthSystemFields

export type CollectionRecords = {
	bookmark_tags: BookmarkTagsRecord
	bookmarks: BookmarksRecord
	collections: CollectionsRecord
	groups: GroupsRecord
	tags: TagsRecord
	users: UsersRecord
}