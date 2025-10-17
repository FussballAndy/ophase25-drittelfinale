export interface Station {
    name: string,
    location: string,
}

export interface Group {
    stations: number[],
}

export interface WinnerSub {
    token: string,
    iteration: number,
    score: number,
}