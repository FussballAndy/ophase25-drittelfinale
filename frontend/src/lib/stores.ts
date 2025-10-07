import { writable } from "svelte/store";
import type { Group, Station } from "./types";

export const pageStore = writable(0);
export const userStore = writable<string | null>(localStorage.getItem("token"));
export const stationsStore: {
    stations: Station[],
    groups: Group[],
} = {
    stations: [{
        name: "Sudoku",
        location: "Zuhause"
    }],
    groups: [{
        points: 0,
        stations: [0, 0, 0]
    },{
        points: 1,
        stations: [0, 0, 0]
    },{
        points: 2,
        stations: [0, 0, 0]
    }]
};