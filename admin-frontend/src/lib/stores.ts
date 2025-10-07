import { writable } from "svelte/store";

export const pageStore = writable<number>(3);
export const requestNameStore = writable<string>("")
export const requestDataStore = writable<any[]>([])

interface Group {
    stations: number[]
}

interface Station {
    name: string,
    location: string,
}

interface Question {
    num: number,
    prompt: string,
    answers: string[]
}

export const groupsStore = writable<Group[]>([])
export const stationsStore = writable<Station[]>([])
export const questionsStore = writable<Question[]>([{
    num: 0,
    prompt: "Abc",
    answers: []
}])
export const tokenStore = writable<Record<string, number>>({})
export const scoreStore = writable<(number | undefined)[][]>(Array(25).fill(Array(4).fill(undefined)))

export const scalingStore = writable(1)
export const drittelReady = writable(true)