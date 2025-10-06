import { writable } from "svelte/store"

let websocket: WebSocket | null = null

export interface Question {
    num: number,
    prompt: string,
    answers: string[]
}

export interface Submission {
    question: number,
    answer: number
}

export let questionStore = writable<Question | null>(null)

export function wsStart() {
    websocket = new WebSocket("ws://localhost:8080/api/drittel")
    websocket.addEventListener("message", e => {
        questionStore.set(JSON.parse(e.data))
    })
}

export function wsClose() {
    websocket?.close()
}

export function wsSubmitAnswer(question: number, answer: number) {
    const submission: Submission = {
        question,
        answer
    }
    websocket?.send(JSON.stringify(submission))
}