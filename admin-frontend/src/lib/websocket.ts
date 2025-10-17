import { get } from "svelte/store";
import RequestPage from "../pages/RequestPage.svelte";
import { config } from "./config";
import { pageStore, requestNameStore, requestDataStore, groupsStore, stationsStore, questionsStore, tokenStore, scoreStore, drittelReady, questionResultStore, questionActiveStore } from "./stores";

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

function handleMessage(msg: JSONServer) {
    console.log(msg)
    switch(msg.kind) {
        case "num_stations":
            config.num_stations = msg.data;
            break;
        case "groups": case "stations": case "questions": case "tokens":
            pageStore.set(1)
            requestNameStore.set(msg.kind)
            if(msg.kind === "tokens") {
                requestDataStore.set([msg.data])
            } else {
                requestDataStore.set(msg.data)
            }
            break
        case "confirmation":
            pageStore.set(2)
            break
        case "ingame":
            pageStore.set(3)
            
            //const scoreData = Array.from(atob(msg.data), char => (char as string).charCodeAt(0));
            const scoreData = msg.data;
            console.log(scoreData)
            const scoreArray = new Array(25).fill([])
            for(let i = 0; i < scoreArray.length; i++) {
                scoreArray[i] = new Array(4).fill(undefined)
                for(let j = 0; j < scoreArray[i].length; j++) {
                    if(scoreData[i*4+j] !== 255) {
                        scoreArray[i][j] = scoreData[i*4+j];
                    }
                }
            }
            console.log(scoreArray)
            scoreStore.set(scoreArray)
            break
        case "result":
            const cScore = get(scoreStore);
            cScore[msg.data.station][msg.data.iter] = msg.data.result
            scoreStore.set(cScore)
            break
        case "final":
            // ignore
            drittelReady.set(true)
            break
        case "front":
            alert("Front status: " + msg.data)
            break
        case "scores":
            break
        case "drittel":
            break
        case "submissions":
            questionResultStore.update(v => {
                v[get(questionActiveStore)!] = msg.data
                return v
            })
            break
        case "end":
            break
        case "fulldump":
            stationsStore.set(msg.data.stations)
            tokenStore.set(msg.data.tokens)
            groupsStore.set(msg.data.groups)
            questionsStore.set(msg.data.questions)
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

function exportData(data: any, name: string) {
    const link = document.createElement('a');
    link.download = name;
    const blob = new Blob([JSON.stringify(data)], {type: 'application/json'});
    link.href = window.URL.createObjectURL(blob);
    link.click();
}