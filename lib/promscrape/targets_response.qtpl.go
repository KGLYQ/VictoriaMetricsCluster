// Code generated by qtc from "targets_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line lib/promscrape/targets_response.qtpl:1
package promscrape

//line lib/promscrape/targets_response.qtpl:1
import "github.com/VictoriaMetrics/VictoriaMetrics/lib/prompbmarshal"

//line lib/promscrape/targets_response.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line lib/promscrape/targets_response.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line lib/promscrape/targets_response.qtpl:6
func StreamTargetsResponsePlain(qw422016 *qt422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, showOriginLabels bool) {
//line lib/promscrape/targets_response.qtpl:8
	for _, js := range jts {
//line lib/promscrape/targets_response.qtpl:8
		qw422016.N().S(`job=`)
//line lib/promscrape/targets_response.qtpl:9
		qw422016.N().Q(js.job)
//line lib/promscrape/targets_response.qtpl:9
		qw422016.N().S(`(`)
//line lib/promscrape/targets_response.qtpl:9
		qw422016.N().D(js.upCount)
//line lib/promscrape/targets_response.qtpl:9
		qw422016.N().S(`/`)
//line lib/promscrape/targets_response.qtpl:9
		qw422016.N().D(js.targetsTotal)
//line lib/promscrape/targets_response.qtpl:9
		qw422016.N().S(`up)`)
//line lib/promscrape/targets_response.qtpl:10
		qw422016.N().S(`
`)
//line lib/promscrape/targets_response.qtpl:11
		for _, ts := range js.targetsStatus {
//line lib/promscrape/targets_response.qtpl:13
			labels := promLabelsString(ts.labels)
			ol := promLabelsString(ts.originalLabels)

//line lib/promscrape/targets_response.qtpl:16
			qw422016.N().S("\t")
//line lib/promscrape/targets_response.qtpl:16
			qw422016.N().S(`state=`)
//line lib/promscrape/targets_response.qtpl:16
			if ts.up {
//line lib/promscrape/targets_response.qtpl:16
				qw422016.N().S(`up`)
//line lib/promscrape/targets_response.qtpl:16
			} else {
//line lib/promscrape/targets_response.qtpl:16
				qw422016.N().S(`down`)
//line lib/promscrape/targets_response.qtpl:16
			}
//line lib/promscrape/targets_response.qtpl:16
			qw422016.N().S(`,`)
//line lib/promscrape/targets_response.qtpl:16
			qw422016.N().S(` `)
//line lib/promscrape/targets_response.qtpl:16
			qw422016.N().S(`endpoint=`)
//line lib/promscrape/targets_response.qtpl:17
			qw422016.N().S(ts.endpoint)
//line lib/promscrape/targets_response.qtpl:17
			qw422016.N().S(`,{ %space %}labels=`)
//line lib/promscrape/targets_response.qtpl:18
			qw422016.N().S(labels)
//line lib/promscrape/targets_response.qtpl:19
			if showOriginLabels {
//line lib/promscrape/targets_response.qtpl:19
				qw422016.N().S(`, originalLabels=`)
//line lib/promscrape/targets_response.qtpl:19
				qw422016.N().S(ol)
//line lib/promscrape/targets_response.qtpl:19
			}
//line lib/promscrape/targets_response.qtpl:19
			qw422016.N().S(`,`)
//line lib/promscrape/targets_response.qtpl:19
			qw422016.N().S(` `)
//line lib/promscrape/targets_response.qtpl:19
			qw422016.N().S(`last_scrape=`)
//line lib/promscrape/targets_response.qtpl:20
			qw422016.N().FPrec(ts.lastScrapeTime.Seconds(), 3)
//line lib/promscrape/targets_response.qtpl:20
			qw422016.N().S(`s ago,`)
//line lib/promscrape/targets_response.qtpl:20
			qw422016.N().S(` `)
//line lib/promscrape/targets_response.qtpl:20
			qw422016.N().S(`scrape_duration=`)
//line lib/promscrape/targets_response.qtpl:21
			qw422016.N().FPrec(ts.scrapeDuration.Seconds(), 3)
//line lib/promscrape/targets_response.qtpl:21
			qw422016.N().S(`s,`)
//line lib/promscrape/targets_response.qtpl:21
			qw422016.N().S(` `)
//line lib/promscrape/targets_response.qtpl:21
			qw422016.N().S(`samples_scraped=`)
//line lib/promscrape/targets_response.qtpl:22
			qw422016.N().D(ts.samplesScraped)
//line lib/promscrape/targets_response.qtpl:22
			qw422016.N().S(`,`)
//line lib/promscrape/targets_response.qtpl:22
			qw422016.N().S(` `)
//line lib/promscrape/targets_response.qtpl:22
			qw422016.N().S(`error=`)
//line lib/promscrape/targets_response.qtpl:23
			qw422016.N().Q(ts.errMsg)
//line lib/promscrape/targets_response.qtpl:24
			qw422016.N().S(`
`)
//line lib/promscrape/targets_response.qtpl:25
		}
//line lib/promscrape/targets_response.qtpl:26
	}
//line lib/promscrape/targets_response.qtpl:28
	for _, jobName := range emptyJobs {
//line lib/promscrape/targets_response.qtpl:28
		qw422016.N().S(`job=`)
//line lib/promscrape/targets_response.qtpl:29
		qw422016.N().Q(jobName)
//line lib/promscrape/targets_response.qtpl:29
		qw422016.N().S(`(0/0 up)`)
//line lib/promscrape/targets_response.qtpl:30
		qw422016.N().S(`
`)
//line lib/promscrape/targets_response.qtpl:31
	}
//line lib/promscrape/targets_response.qtpl:33
}

//line lib/promscrape/targets_response.qtpl:33
func WriteTargetsResponsePlain(qq422016 qtio422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, showOriginLabels bool) {
//line lib/promscrape/targets_response.qtpl:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/targets_response.qtpl:33
	StreamTargetsResponsePlain(qw422016, jts, emptyJobs, showOriginLabels)
//line lib/promscrape/targets_response.qtpl:33
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/targets_response.qtpl:33
}

//line lib/promscrape/targets_response.qtpl:33
func TargetsResponsePlain(jts []jobTargetsStatuses, emptyJobs []string, showOriginLabels bool) string {
//line lib/promscrape/targets_response.qtpl:33
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/targets_response.qtpl:33
	WriteTargetsResponsePlain(qb422016, jts, emptyJobs, showOriginLabels)
//line lib/promscrape/targets_response.qtpl:33
	qs422016 := string(qb422016.B)
//line lib/promscrape/targets_response.qtpl:33
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/targets_response.qtpl:33
	return qs422016
//line lib/promscrape/targets_response.qtpl:33
}

//line lib/promscrape/targets_response.qtpl:35
func StreamTargetsResponseHTML(qw422016 *qt422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, redirectPath string, onlyUnhealthy bool) {
//line lib/promscrape/targets_response.qtpl:35
	qw422016.N().S(`<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-giJF6kkoqNQ00vy+HMDP7azOuL0xtbfIcaT9wjKHr8RbDVddVHyTfAAsrekwKmP1" crossorigin="anonymous"><title>Scrape targets</title></head><body class="m-3"><h1>Scrape targets</h1><div><button type="button" class="btn`)
//line lib/promscrape/targets_response.qtpl:47
	if !onlyUnhealthy {
//line lib/promscrape/targets_response.qtpl:47
		qw422016.N().S(`btn-primary`)
//line lib/promscrape/targets_response.qtpl:47
	} else {
//line lib/promscrape/targets_response.qtpl:47
		qw422016.N().S(`btn-secondary`)
//line lib/promscrape/targets_response.qtpl:47
	}
//line lib/promscrape/targets_response.qtpl:47
	qw422016.N().S(`"`)
//line lib/promscrape/targets_response.qtpl:48
	if onlyUnhealthy {
//line lib/promscrape/targets_response.qtpl:48
		qw422016.N().S(`onclick="location.href='`)
//line lib/promscrape/targets_response.qtpl:48
		qw422016.E().S(redirectPath)
//line lib/promscrape/targets_response.qtpl:48
		qw422016.N().S(`'"`)
//line lib/promscrape/targets_response.qtpl:48
	}
//line lib/promscrape/targets_response.qtpl:48
	qw422016.N().S(`>All</button><button type="button" class="btn`)
//line lib/promscrape/targets_response.qtpl:51
	if onlyUnhealthy {
//line lib/promscrape/targets_response.qtpl:51
		qw422016.N().S(`btn-primary`)
//line lib/promscrape/targets_response.qtpl:51
	} else {
//line lib/promscrape/targets_response.qtpl:51
		qw422016.N().S(`btn-secondary`)
//line lib/promscrape/targets_response.qtpl:51
	}
//line lib/promscrape/targets_response.qtpl:51
	qw422016.N().S(`"`)
//line lib/promscrape/targets_response.qtpl:52
	if !onlyUnhealthy {
//line lib/promscrape/targets_response.qtpl:52
		qw422016.N().S(`onclick="location.href='`)
//line lib/promscrape/targets_response.qtpl:52
		qw422016.N().S(redirectPath)
//line lib/promscrape/targets_response.qtpl:52
		qw422016.N().S(`?show_only_unhealthy=true'"`)
//line lib/promscrape/targets_response.qtpl:52
	}
//line lib/promscrape/targets_response.qtpl:52
	qw422016.N().S(`>Unhealthy</button></div>`)
//line lib/promscrape/targets_response.qtpl:56
	for _, js := range jts {
//line lib/promscrape/targets_response.qtpl:57
		if onlyUnhealthy && js.upCount == js.targetsTotal {
//line lib/promscrape/targets_response.qtpl:57
			continue
//line lib/promscrape/targets_response.qtpl:57
		}
//line lib/promscrape/targets_response.qtpl:57
		qw422016.N().S(`<div><h4><a>`)
//line lib/promscrape/targets_response.qtpl:60
		qw422016.E().S(js.job)
//line lib/promscrape/targets_response.qtpl:60
		qw422016.N().S(`(`)
//line lib/promscrape/targets_response.qtpl:60
		qw422016.N().D(js.upCount)
//line lib/promscrape/targets_response.qtpl:60
		qw422016.N().S(`/`)
//line lib/promscrape/targets_response.qtpl:60
		qw422016.N().D(js.targetsTotal)
//line lib/promscrape/targets_response.qtpl:60
		qw422016.N().S(`up)</a></h4><table class="table table-striped table-hover table-bordered table-sm"><thead><tr><th scope="col">Endpoint</th><th scope="col">State</th><th scope="col">Labels</th><th scope="col">Last Scrape</th><th scope="col">Scrape Duration</th><th scope="col">Samples Scraped</th><th scope="col">Error</th></tr></thead><tbody>`)
//line lib/promscrape/targets_response.qtpl:75
		for _, ts := range js.targetsStatus {
//line lib/promscrape/targets_response.qtpl:76
			if onlyUnhealthy && ts.up {
//line lib/promscrape/targets_response.qtpl:76
				continue
//line lib/promscrape/targets_response.qtpl:76
			}
//line lib/promscrape/targets_response.qtpl:76
			qw422016.N().S(`<tr`)
//line lib/promscrape/targets_response.qtpl:77
			if !ts.up {
//line lib/promscrape/targets_response.qtpl:77
				qw422016.N().S(`class="alert alert-danger" role="alert"`)
//line lib/promscrape/targets_response.qtpl:77
			}
//line lib/promscrape/targets_response.qtpl:77
			qw422016.N().S(`><td><a href="`)
//line lib/promscrape/targets_response.qtpl:78
			qw422016.E().S(ts.endpoint)
//line lib/promscrape/targets_response.qtpl:78
			qw422016.N().S(`">`)
//line lib/promscrape/targets_response.qtpl:78
			qw422016.E().S(ts.endpoint)
//line lib/promscrape/targets_response.qtpl:78
			qw422016.N().S(`</a><br></td><td>`)
//line lib/promscrape/targets_response.qtpl:79
			if ts.up {
//line lib/promscrape/targets_response.qtpl:79
				qw422016.N().S(`UP`)
//line lib/promscrape/targets_response.qtpl:79
			} else {
//line lib/promscrape/targets_response.qtpl:79
				qw422016.N().S(`DOWN`)
//line lib/promscrape/targets_response.qtpl:79
			}
//line lib/promscrape/targets_response.qtpl:79
			qw422016.N().S(`</td><td title="Original labels:`)
//line lib/promscrape/targets_response.qtpl:80
			streamformatLabel(qw422016, ts.originalLabels)
//line lib/promscrape/targets_response.qtpl:80
			qw422016.N().S(`">`)
//line lib/promscrape/targets_response.qtpl:81
			streamformatLabel(qw422016, ts.labels)
//line lib/promscrape/targets_response.qtpl:81
			qw422016.N().S(`</td><td>`)
//line lib/promscrape/targets_response.qtpl:83
			qw422016.N().FPrec(ts.lastScrapeTime.Seconds(), 3)
//line lib/promscrape/targets_response.qtpl:83
			qw422016.N().S(`s ago</td><td>`)
//line lib/promscrape/targets_response.qtpl:84
			qw422016.N().FPrec(ts.scrapeDuration.Seconds(), 3)
//line lib/promscrape/targets_response.qtpl:84
			qw422016.N().S(`s</td><td>`)
//line lib/promscrape/targets_response.qtpl:85
			qw422016.N().D(ts.samplesScraped)
//line lib/promscrape/targets_response.qtpl:85
			qw422016.N().S(`</td><td>`)
//line lib/promscrape/targets_response.qtpl:86
			qw422016.E().S(ts.errMsg)
//line lib/promscrape/targets_response.qtpl:86
			qw422016.N().S(`</td></tr>`)
//line lib/promscrape/targets_response.qtpl:88
		}
//line lib/promscrape/targets_response.qtpl:88
		qw422016.N().S(`</tbody></table></div>`)
//line lib/promscrape/targets_response.qtpl:92
	}
//line lib/promscrape/targets_response.qtpl:94
	for _, jobName := range emptyJobs {
//line lib/promscrape/targets_response.qtpl:94
		qw422016.N().S(`<div><h4><a>`)
//line lib/promscrape/targets_response.qtpl:97
		qw422016.E().S(jobName)
//line lib/promscrape/targets_response.qtpl:97
		qw422016.N().S(`(0/0 up)</a></h4><table class="table table-striped table-hover table-bordered table-sm"><thead><tr><th scope="col">Endpoint</th><th scope="col">State</th><th scope="col">Labels</th><th scope="col">Last Scrape</th><th scope="col">Scrape Duration</th><th scope="col">Samples Scraped</th><th scope="col">Error</th></tr></thead></table></div>`)
//line lib/promscrape/targets_response.qtpl:113
	}
//line lib/promscrape/targets_response.qtpl:113
	qw422016.N().S(`</body></html>`)
//line lib/promscrape/targets_response.qtpl:116
}

//line lib/promscrape/targets_response.qtpl:116
func WriteTargetsResponseHTML(qq422016 qtio422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, redirectPath string, onlyUnhealthy bool) {
//line lib/promscrape/targets_response.qtpl:116
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/targets_response.qtpl:116
	StreamTargetsResponseHTML(qw422016, jts, emptyJobs, redirectPath, onlyUnhealthy)
//line lib/promscrape/targets_response.qtpl:116
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/targets_response.qtpl:116
}

//line lib/promscrape/targets_response.qtpl:116
func TargetsResponseHTML(jts []jobTargetsStatuses, emptyJobs []string, redirectPath string, onlyUnhealthy bool) string {
//line lib/promscrape/targets_response.qtpl:116
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/targets_response.qtpl:116
	WriteTargetsResponseHTML(qb422016, jts, emptyJobs, redirectPath, onlyUnhealthy)
//line lib/promscrape/targets_response.qtpl:116
	qs422016 := string(qb422016.B)
//line lib/promscrape/targets_response.qtpl:116
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/targets_response.qtpl:116
	return qs422016
//line lib/promscrape/targets_response.qtpl:116
}

//line lib/promscrape/targets_response.qtpl:118
func streamformatLabel(qw422016 *qt422016.Writer, labels []prompbmarshal.Label) {
//line lib/promscrape/targets_response.qtpl:119
	for _, label := range labels {
//line lib/promscrape/targets_response.qtpl:120
		qw422016.E().S(label.Name)
//line lib/promscrape/targets_response.qtpl:120
		qw422016.N().S(`=`)
//line lib/promscrape/targets_response.qtpl:120
		qw422016.E().Q(label.Value)
//line lib/promscrape/targets_response.qtpl:120
		qw422016.N().S(` `)
//line lib/promscrape/targets_response.qtpl:121
	}
//line lib/promscrape/targets_response.qtpl:122
}

//line lib/promscrape/targets_response.qtpl:122
func writeformatLabel(qq422016 qtio422016.Writer, labels []prompbmarshal.Label) {
//line lib/promscrape/targets_response.qtpl:122
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/targets_response.qtpl:122
	streamformatLabel(qw422016, labels)
//line lib/promscrape/targets_response.qtpl:122
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/targets_response.qtpl:122
}

//line lib/promscrape/targets_response.qtpl:122
func formatLabel(labels []prompbmarshal.Label) string {
//line lib/promscrape/targets_response.qtpl:122
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/targets_response.qtpl:122
	writeformatLabel(qb422016, labels)
//line lib/promscrape/targets_response.qtpl:122
	qs422016 := string(qb422016.B)
//line lib/promscrape/targets_response.qtpl:122
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/targets_response.qtpl:122
	return qs422016
//line lib/promscrape/targets_response.qtpl:122
}
