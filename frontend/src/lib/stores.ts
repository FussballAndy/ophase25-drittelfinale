import { writable } from "svelte/store";
import type { Group, Station } from "./types";
import { API_BASE_URL } from "./consts";

export const pageStore = writable(0);
export const userStore = writable<string | null>(localStorage.getItem("token"));
const stationStorage = localStorage.getItem("station");
const stationString = stationStorage === null ? -1 : parseInt(stationStorage)
export const userStationStore = writable(stationString)
export const stationsStore = writable<Station[]>([])
export const groupsStore = writable<Group[]>([])

fetch(API_BASE_URL + "/api/stations").then(res => res.json()).then(res => stationsStore.set(res))
fetch(API_BASE_URL + "/api/groups").then(res => res.json()).then(res => groupsStore.set(res))