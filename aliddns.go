package main

import (
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

func getRRByHostname() string {
	hostname, _ := os.Hostname()
	hnList := strings.Split(hostname, ".")
	if len(hnList) == 0 {
		return ""
	}
	return hnList[0]
}

func localIP() net.IP {
	ifaddrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Printf("InterfaceAddrs error: %+v\n", err)
	}
	for _, addr := range ifaddrs {
		ip := net.ParseIP(strings.Split(addr.String(), "/")[0])
		if ip != nil && ip.To4() == nil && ip.To16().IsGlobalUnicast() {
			return ip
		}
	}
	log.Printf("no available global unicast ipv6: %+v\n", net.IPv6zero)
	return net.IPv6zero
}

func getRecordList(client *alidns.Client) []alidns.Record {
	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = GConfig.Domain
	resp, err := client.DescribeDomainRecords(request)
	if err != nil {
		log.Printf("DescribeDomainRecords error: %+v\n", err)
	}
	return resp.DomainRecords.Record
}
func addRecord(client *alidns.Client) (*alidns.AddDomainRecordResponse, error) {
	request := alidns.CreateAddDomainRecordRequest()
	request.DomainName = GConfig.Domain
	request.RR = getRRByHostname()
	request.Type = "AAAA"
	request.Value = localIP().String()
	log.Printf("need to be added: %+v\t%+v\n", request.RR+"."+GConfig.Domain, localIP().String())
	return client.AddDomainRecord(request)
}
func updateRecord(client *alidns.Client, recordID string) (*alidns.UpdateDomainRecordResponse, error) {
	request := alidns.CreateUpdateDomainRecordRequest()
	request.RecordId = recordID
	request.RR = getRRByHostname()
	request.Type = "AAAA"
	request.Value = localIP().String()
	log.Printf("need to be updated: %+v\t%+v\n", request.RR+"."+GConfig.Domain, localIP().String())
	return client.UpdateDomainRecord(request)
}
func deleteRecord(client *alidns.Client, recordID string) (*alidns.DeleteDomainRecordResponse, error) {
	request := alidns.CreateDeleteDomainRecordRequest()
	request.RecordId = recordID
	log.Printf("need to be deleted:%+v\n", getRRByHostname()+"."+GConfig.Domain)
	return client.DeleteDomainRecord(request)
}

func dnsSync() {
	client, err := alidns.NewClientWithAccessKey("default", GConfig.AliDNS.ApiKey, GConfig.AliDNS.ApiSecret)
	if err != nil {
		log.Printf("%+v\n", err)
		return
	}
	var matchCount int
	for _, v := range getRecordList(client) {
		if v.Type == "AAAA" && v.RR == getRRByHostname() {
			log.Printf("current record: %+v\t%+v\n", v.RR+"."+GConfig.Domain, v.Value)
			if matchCount == 0 {
				if v.Value != localIP().String() && !localIP().Equal(net.IPv6zero) {
					resp, err := updateRecord(client, v.RecordId)
					if err != nil {
						log.Printf("updateRecord error: %+v\n", err)
					}
					log.Printf("record updated: %+v\n", resp)
				}
			} else {
				resp, err := deleteRecord(client, v.RecordId)
				if err != nil {
					log.Printf("deleteRecord error: %+v\n", err)
				}
				log.Printf("record deleted: %+v\n", resp)
			}
			matchCount++
		}
	}
	if !localIP().Equal(net.IPv6zero) && matchCount == 0 {
		resp, err := addRecord(client)
		if err != nil {
			log.Printf("addRecord error: %+v\n", err)
		}
		log.Printf("record added: %+v\n", resp)
	}
}

func aliddns() {
	dnsSync()
	tk := time.NewTicker(time.Duration(GConfig.Interval) * time.Second)
	defer tk.Stop()
	for range tk.C {
		dnsSync()
	}
}
