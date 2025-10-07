import { get } from "svelte/store";
import RequestPage from "../pages/RequestPage.svelte";
import { config } from "./config";
import { pageStore, requestNameStore, requestDataStore, groupsStore, stationsStore, questionsStore, tokenStore, scoreStore, drittelReady } from "./stores";

let websocket: WebSocket | undefined;

export function wsStart(token: string) {
    websocket = new WebSocket("http://127.0.0.1:8080/api/admin?session=" + token)
    websocket.addEventListener("message", e => handleMessage(JSON.parse(e.data)))
}

export function wsClose() {
    websocket && websocket.close()
}

interface JSONServer {
    kind: string,
    data: any,
}

function handleMessage(data: JSONServer) {
    console.log(data)
    switch(data.kind) {
        case "num_stations":
            config.num_stations = data.data;
            break;
        case "groups": case "stations": case "questions": case "tokens":
            pageStore.set(1)
            requestNameStore.set(data.kind)
            if(data.kind == "tokens") {
                requestDataStore.set([data.data])
            } else {
                requestDataStore.set(data.data)
            }
            break
        case "confirmation":
            pageStore.set(2)
            break
        case "ingame":
            pageStore.set(3)
            break
        case "result":
            scoreStore.update(score => {
                score[data.data.station][data.data.iter] = data.data.result
                return score
            })
            break
        case "final":
            // ignore
            drittelReady.set(true)
            break
    }
}

export function wsSend(data: any) {
    if(websocket) {
        switch(get(requestNameStore)) {
            case "groups":
                groupsStore.set(data)
            case "stations":
                stationsStore.set(data)
            case "questions":
                questionsStore.set(data)
            case "tokens":
                tokenStore.set(data[0])
        }
        websocket.send(JSON.stringify(data))
        console.log("Sent ", JSON.stringify(data))
    }
}