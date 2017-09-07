// Code generated by protoc-gen-js-fetch.
// DO NOT EDIT!

import * as apipb_messages from "./messages.pb";
import * as apipb_models from "./models.pb";

export type CategoryReqTask =  "ADD"  | "UPDATE"  | "DELETE" ;
export const CategoryReqTask_ADD: CategoryReqTask = "ADD";
export const CategoryReqTask_UPDATE: CategoryReqTask = "UPDATE";
export const CategoryReqTask_DELETE: CategoryReqTask = "DELETE";

export class AllGameReleases {
  gameId: string;
  releases: AllGameReleasesMiniRelease[];
}

export class AllGameReleasesMiniRelease {
  version: string;
  releasedAt: string|number;
  releaseState: apipb_models.ReleaseState;
}

export class GameCategoryLocale {
  language: string;
  title: string;
  image: string;
  color: string;
  description: string;
}

export class GameCategory {
  name: string;
  locales: GameCategoryLocale[];
  revision: number;
  labels: { [key: string]: string };
		  defaultLocale: string;
}

export class CategoryReq {
  task: CategoryReqTask;
  category: GameCategory;
  knownRevision: number;
}

export class CategoryListReq {
}

export class CategoryList {
  categories: GameCategory[];
}

export class GetAllGamesReq {
/**
Games that user wants the data of, if it is empty returns all games
*/
  games: apipb_messages.GameAndVersion[];
/**
Language filters games. If language field is empty than returns games with all languages.
*/
  language: string;
}

export class GetAllGamesRes {
  games: apipb_models.GameRelease[];
}

export class PublishReq {
  manifest: apipb_models.GameManifest;
  files: { [key: string]: string };
		}

export class PublishRes {
  token: apipb_models.UploadToken;
  uploadUrls: { [key: string]: string };
		  storage: string;
}

export class TarballInfo {
  url: string;
  storage: string;
  archiveFormat: string;
}

export class AddTarballReq {
  token: string;
  infos: TarballInfo[];
}

export class AddTarballRes {
  packageUrls: { [key: string]: string };
		}

export class RegistryClient {
  uniqueName: string;
  apiKey: string;
}

export class CreateClientReq {
  uniqueName: string;
}

export class RevokeClientReq {
  apiKey: string;
}

export class ClientList {
  clientNames: string[];
}

export class ListClientReq {
}
