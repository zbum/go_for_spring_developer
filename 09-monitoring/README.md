## 스카우터 
* 다음은 golang으로 구현한 응용시스템을 스카우터 APM으로 모니터링하기 위해 개발자가 작업해야 하는 내용을 정리한 가이드입니다.
* golang 은 java 의 ASM과 같은 기술이 없기(?) 때문에 침습적(invasive)으로 agent 를 적용해야 합니다.
* 최대한 비즈니스 로직 구현 코드에 기술코드가 섞이지 않도록 미들웨어(middleware), 플러그인을 추가하였습니다.

## 설치
### 라이브러리
* go.mod 파일에 다음을 추가하여 library 를 설정합니다.
* 스카우터가 제공하는 golang agent 가 더 이상 업데이트 되지 않고 있어 fork 를 생성하여 필요한 기능을 추가하였습니다.
    * https://github.com/zbum/scouter-agent-golang

```
go get github.com/zbum/scouter-agent-golang@v0.20.1-alpha.1
````

## 개발 가이드
### main.go 설정
* main 함수에서 scouterx.Init() 으로 스카우터를 초기화 합니다.
```go
 
package main

import (
	"github.com/zbum/scouter-agent-golang/scouterx"
	"golang-restapi-prototype/logger"
	"golang-restapi-prototype/route"
)

var log = logger.GetLogger()

func main() {
	scouterx.Init()
	route.HandleRequest()
	var _ = log.Sync()
}

```

### middleware 설정
* fiber 를 사용한다면 다음과 같이 middleware 를 추가합니다.

```go

import (
	scouter_middleware "github.com/zbum/scouter-agent-golang/scouterx/middleware"
)

app := fiber.New(fiber.Config{
    Views:                        engine,
    JSONEncoder:                  json.Marshal,
    JSONDecoder:                  json.Unmarshal,
    DisablePreParseMultipartForm: true,
    StreamRequestBody:            true,
})

app.Use(
    scouter_middleware.FiberTracingMiddleware()
)
```

* net/http 기반 프레임워크(ex. gorilla mux) 를 사용한다면 다음과 같이 middleware 를 추가합니다.

```go

import (
	scouter_middleware "github.com/zbum/scouter-agent-golang/scouterx/middleware"
)

r := mux.NewRouter().StrictSlash(true)
r.Use( scouter_middleware.HttpTracingMiddleware )
```

### function profile
* 함수의 profile 을 확인하기 위해서는 함수 시작에 다음의 코드를 추가해야 합니다.

```go
ctx := r.Context() // Context 를 구해옵니다.
step := strace.StartMethod(ctx, methodName)
defer strace.EndMethod(ctx, step)
```

### api call profile
* api 호출의 profile 을 위한 함수도 스카우터에서 제공합니다.
* HttpClient 객체를 초기화 할때 Transport 를 ScouterRoundTripper 로 Wrapping 하여 api call 을 프로파일링 할 수 있습니다.
```go
var HttpClient = &http.Client{
	Transport: strace.ScouterRoundTripper{Proxied: http.DefaultTransport},
	Timeout:   5 * time.Second,
}
```

### sql profile
* gorm 을 사용하는 경우, gorm.DB 를 초기화 하는 곳에서 strace.GormDbPlugin{} 을 Plugin 으로 등록합니다.
* 이후, 특정 이벤트에서 코드를 추가할 필요는 없습니다.
```go
db, err := gorm.Open(mysql.Open(g.dsn()), &gorm.Config{
    Logger: g.logger,
})
if err != nil {
    g.logger.Error(context.Background(), "error %s during the open db\n", err)
    return err
}

err = db.Use(strace.GormDbPlugin{})
```

* 다만, gorm 을 사용하여 sql 을 실행할때, WithContext 함수를 활용하여  context.Context 를 포함하여 실행해야 합니다.
```
func (ur *DefaultUserRepository) GetUser(ctx *context.Context, q *query.Query, userId int64) (*domain.User, error) {
	step := strace.StartMethod(*ctx)
	defer strace.EndMethod(*ctx, step)

	d := q.Department
	u := q.User

	user, err := u.
		WithContext(*ctx).
		Select(u.ALL, d.ALL).
		LeftJoin(d, d.DepartmentId.EqCol(u.DepartmentId)).
		Where(u.Id.Eq(userId)).
		Take()

	return user, err
}
```


## 실행
* 스카우터 실행시 스카우터 설정파일의 위치를 OS 환경 변수(SCOUTER_CONFIG)로 제공해야 합니다.
* 설정파일의 예 (scouter.conf)
```
obj_name=gfsd
obj_host_name=gfsd1

# Scouter Server IP Address (Default : 127.0.0.1)
net_collector_ip=127.0.0.1

# Scouter Server Port (Default : 6100)
net_collector_udp_port=6100
net_collector_tcp_port=6100

profile_http_querystring_enabled=true
profile_http_header_enabled=true
```



## 결과
* 다음과 같이 메소드 프로파일이 적용되었는지 확인합니다.
![img.png](img.png)

## Prometheus

* prometheus api 는 promhttp.Handler() 를 mux.Router 에 등록하여 사용합니다.
```go
package handler

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func AddPrometheusHandler(r *mux.Router) {
	r.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)
}
```

* 설정 후 조회하면 다음의 결과를 확인 할 수 있습니다.
```go
curl http://localhost:15510/metrics
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 8
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.19.5"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 1.072064e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 1.072064e+06
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 4642
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 711
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 3.684152e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 1.072064e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 4.349952e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 3.547136e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 8231
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 4.349952e+06
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 7.897088e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 8942
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 12000
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 15600
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 67392
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 81360
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 1.11615e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 491520
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 491520
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 1.3290512e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 8
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 3
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
➜  golang-restapi-prototype git:(main) ✗ curl http://localhost:15510/metrics
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 9
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.19.5"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 1.113008e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 1.113008e+06
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 4642
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 711
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 3.684152e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 1.113008e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 4.284416e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 3.612672e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 8306
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 4.284416e+06
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 7.897088e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 9017
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 12000
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 15600
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 67392
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 81360
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 1.11615e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 491520
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 491520
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 1.3290512e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 8
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 4
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```