#include <linux/bpf.h>
#include <linux/pkt_cls.h>
#include <linux/ip.h>
#include <linux/if_ether.h>
#include <bpf/bpf_helpers.h>
#include <netinet/in.h>

SEC("out_block_c2") 
int egress(struct __sk_buff *skb) {   

    // Direct Packet Access from skbuff
    void* packet_data = (void*)(long)skb->data;
    void* packet_end = (void*)(long)skb->data_end;
    
    struct iphdr*    packet_ip_header;
    struct ethhdr*   packet_eth_header;
    
    // JIT Verifier - Check we haven't gone beyond end of packet
    if(packet_data + sizeof(*packet_eth_header) > packet_end)
        return TC_ACT_OK;

    // JIT Verifier - Check the Ethernet + IP Frame isn't beyond the length of the packet
    if(packet_data + sizeof(*packet_eth_header) + sizeof(*packet_ip_header) > packet_end)
        return TC_ACT_OK;
    
    // Start of ip header = [start of packet]+[Ethernet Frame length] since we are looking
    // for an IP packet encapsulated in an ethernet frame.
    packet_ip_header = packet_data + sizeof(*packet_eth_header);
    
    // Discard any non UDP/TCP packets, we don't use them but it has the benefit of
    // ICMP still working for any malware payloads trying to phone home.
    if(     packet_ip_header->protocol == IPPROTO_TCP ||
            packet_ip_header->protocol == IPPROTO_UDP) {
        char fmt_str[] = "packet from:%0x to:%0x using %s";
        bpf_trace_printk(fmt_str, sizeof(fmt_str), 
                packet_ip_header->saddr, 
                packet_ip_header->daddr, 
                packet_ip_header->protocol
        );
        return TC_ACT_OK;
    }

    return TC_ACT_OK;

    //switch (packet_ip_header->protocol) {
    //    case IPPROTO_TCP: case IPPROTO_UDP:
    //        return bpf_map;
    //    default: return TC_ACT_OK;
    //}
}

char __license[] SEC("license") = "Dual MIT/GPL";
