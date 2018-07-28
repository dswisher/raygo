# RayGo

This repo contains my attempt to implement a ray tracer in Go, based on the books by [Peter Shirley](https://en.wikipedia.org/wiki/Peter_Shirley):

* [Ray Tracing in One Weekend](https://www.amazon.com/dp/B01B5AODD8/) - RTIOW - also see this [blog post](http://in1weekend.blogspot.com/2016/01/ray-tracing-in-one-weekend.html), which contains ideas for extending the code
* [Ray Tracing: the Next Week](https://www.amazon.com/dp/B01CO7PQ8C/)
* [Ray Tracing: the Rest of Your Life](https://www.amazon.com/dp/B01DN58P8C/)

The implementations in the book are in C++.
I haven't written C++ in years, so rather than brush up on C++, I decided to try a new language: Go.

I am not the first to do this. See the "Inspiration" section below for links.
Many of them set a high bar, and if you are looking for a *good* example, you are probably better off looking in that list.

## Inspiration

* [pt](https://github.com/fogleman/pt) - a path tracer written in Go by Michael Fogleman. Full featured with lots of examples
* [go-trace](https://github.com/markphelps/go-trace) - RTIOW repo by Mark Phelps. He also blogged about his journey: [Part1](https://markphelps.me/2016/03/15/writing-a-ray-tracer-in-go/) - [Part2](https://markphelps.me/2016/03/26/writing-a-ray-tracer-in-go-2/) - [Part3](https://markphelps.me/2016/05/08/writing-a-ray-tracer-in-go-3/) - [Part4](https://markphelps.me/2016/06/03/writing-a-ray-tracer-in-go-4/) - [Part5](https://markphelps.me/2016/07/24/writing-a-ray-tracer-in-go-5/)
* [ray-tracing](https://github.com/ypujante/ray-tracing) - RTIOW repo by Yan Pujante. He extended the code with parallelism and a preview window using SDL.
* [GoTracer](https://github.com/DheerendraRathor/GoTracer) - RTIOW repo by Dheerendra Rathor. He extended the code with a master/agent distributed rendering mechanism.

## Go Links

* golang.org: [home](https://golang.org/), [docs](https://golang.org/doc/), [packages](https://golang.org/pkg/) and [Effective Go](https://golang.org/doc/effective_go.html)
* [Awesome Go](https://github.com/avelino/awesome-go) - A curated list of awesome Go frameworks, libraries and software.
* [project-layout](https://github.com/golang-standards/project-layout) - a sample project layout with documentation. Also see other large projects written in Go, such as [kubernetes](https://github.com/kubernetes/kubernetes), [CockroachDB](https://github.com/cockroachdb/cockroach) and [Hugo](https://github.com/gohugoio/hugo), or scan the [go topic](https://github.com/topics/go) on github.

