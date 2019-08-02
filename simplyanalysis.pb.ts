// Code generated by protoc-gen-js-fetch.
// DO NOT EDIT!

import * as apipb_datasetmodels from "./datasetmodels.pb";

export type DataInterval =  "HOURLY"  | "DAILY" ;
export const DataInterval_HOURLY: DataInterval = "HOURLY";
export const DataInterval_DAILY: DataInterval = "DAILY";

export const ALL_DataInterval_VALUES: DataInterval[] = [DataInterval_HOURLY,DataInterval_DAILY];

export interface GamePlayingRequest {
  games?: string[];
/**
From is the unix seconds time
*/
  from?: string|number;
/**
To is the unix seconds time
*/
  to?: string|number;
  interval?: DataInterval;
/**
Child Id to filter data from
optional
*/
  childId?: string;
}

export const GamePlayingRequest_games = "games";
export const GamePlayingRequest_from = "from";
export const GamePlayingRequest_to = "to";
export const GamePlayingRequest_interval = "interval";
export const GamePlayingRequest_childId = "child_id";
export interface GamePlayingReply {
  data?: apipb_datasetmodels.DataSet;
}

export const GamePlayingReply_data = "data";
export interface SimplifiedReq {
  analysis?: string;
/**
From is the unix seconds time
*/
  from?: string|number;
/**
To is the unix seconds time
*/
  to?: string|number;
  games?: string[];
  interval?: DataInterval;
/**
Child Id to filter data from
optional
*/
  childId?: string;
}

export const SimplifiedReq_analysis = "analysis";
export const SimplifiedReq_from = "from";
export const SimplifiedReq_to = "to";
export const SimplifiedReq_games = "games";
export const SimplifiedReq_interval = "interval";
export const SimplifiedReq_childId = "child_id";
