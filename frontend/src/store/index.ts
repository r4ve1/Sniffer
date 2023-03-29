import {reactive} from 'vue'

export interface Packet {
    No: number
    Timestamp: number
    Length: number
    Source: string
    Destination: string
    Protocol: string
}

interface Store {
    packets: Packet[]
}

export const store: Store = reactive({
    packets: []
})
