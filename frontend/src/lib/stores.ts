import { writable } from "svelte/store";
import type { Station } from "./types";

export const pageStore = writable(0);
export const userStore = writable<string | null>(localStorage.getItem("token"));
export const stationsStore: {
    stations: Station[]
} = {
    stations: [{
        name: "Sudoku",
        location: "Zuhause"
    }]
};