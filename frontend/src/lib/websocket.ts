import { writable } from "svelte/store"
import { API_BASE_URL } from "./consts"

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
    websocket = new WebSocket(API_BASE_URL + "/api/drittel")
    websocket.addEventListener("message", e => {
        const question = JSON.parse(e.data)
        if(question.num === 255) {
            questionStore.set(null)
        } else {
            questionStore.set(question)
        }
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