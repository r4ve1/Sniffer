import {reactive} from 'vue'

export interface Brief {
    No: number
    Timestamp: number
    Length: number
    Source: string
    Destination: string
    Protocol: string
    Info: string
    Phony: boolean
}

interface ip {
    Version: number
    SrcIP: string
    DstIP: string
}

interface ethernet {
    SrcMAC: string
    DstMAC: string
    EthernetType: string
}

interface tcp {
    SrcPort: string
    DstPort: string
    Seq: number
}

export interface Detail {
    Ip: ip
    Ethernet: ethernet
    Tcp: tcp
}

interface Store {
    showDetail: boolean
    briefs: Brief[]
}

export const store: Store = reactive({
    briefs: [],
    showDetail: false,
})
