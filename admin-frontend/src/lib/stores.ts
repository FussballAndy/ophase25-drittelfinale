import { get, writable } from "svelte/store";

export const pageStore = writable<number>(0);
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
    prompt: string,
    answers: string[],
    correct: number[],
}

export interface SubmissionResult {
    answer: number,
    group: boolean,
    front: boolean,
}

export const groupsStore = writable<Group[]>([])
export const stationsStore = writable<Station[]>([])
export const questionsStore = writable<Question[]>([])
export const tokenStore = writable<Record<string, number>>({})
export const scoreStore = writable<(number | undefined)[][]>(Array(25).fill(Array(4).fill(undefined)))

export const scalingStore = writable(1)
export const drittelReady = writable(true)
export const questionActiveStore = writable<number | null>(null)
export const questionResultStore = writable<{[key: string]: SubmissionResult}[]>(Array(24).fill({}))