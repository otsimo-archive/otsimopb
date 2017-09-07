// Code generated by protoc-gen-js-fetch.
// DO NOT EDIT!


export type ContentListRequestListStatus =  "BOTH"  | "ONLY_DRAFT"  | "ONLY_APPROVED" ;
export const ContentListRequestListStatus_BOTH: ContentListRequestListStatus = "BOTH";
export const ContentListRequestListStatus_ONLY_DRAFT: ContentListRequestListStatus = "ONLY_DRAFT";
export const ContentListRequestListStatus_ONLY_APPROVED: ContentListRequestListStatus = "ONLY_APPROVED";

export type ContentListRequestSortBy =  "WEIGHT"  | "TIME" ;
export const ContentListRequestSortBy_WEIGHT: ContentListRequestSortBy = "WEIGHT";
export const ContentListRequestSortBy_TIME: ContentListRequestSortBy = "TIME";

export type ContentListRequestSortOrder =  "DSC"  | "ASC" ;
export const ContentListRequestSortOrder_DSC: ContentListRequestSortOrder = "DSC";
export const ContentListRequestSortOrder_ASC: ContentListRequestSortOrder = "ASC";

export class Content {
  slug: string;
  title: string;
  language: string;
  date: string|number;
  draft: boolean;
  writtenAt: string;
  author: string;
  category: string;
  url: string;
  weight: number;
  keywords: string[];
  categoryWeight: number;
  markdown: string;
  params: { [key: string]: string };
		}

export class ContentListRequest {
  status: ContentListRequestListStatus;
  limit: number;
  category: string;
  offset: number;
  language: string;
  onlyHtmlUrl: boolean;
  sort: ContentListRequestSortBy;
  order: ContentListRequestSortOrder;
  profileId: string;
  clientVersion: string;
  categories: string[];
  exceptCategories: string[];
}

export class ContentListResponse {
  contents: Content[];
  assetVersion: number;
}

export class ContentGetRequest {
  slug: string;
}
