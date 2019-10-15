package main

import (
	"fmt"

	"github.com/golang/geo/s1"
	"github.com/golang/geo/s2"
)

func main() {
	latlng1 := s2.LatLngFromDegrees(32.78655521159346, -96.79948300123215)
	latlng2 := s2.LatLngFromDegrees(32.78630040569697,-96.79938912391663)
	latlng3 := s2.LatLngFromDegrees(32.78639060255794,-96.79920941591263)
	latlng4 := s2.LatLngFromDegrees(32.786580015668314,-96.79924964904785)

	p1 := s2.PointFromLatLng(latlng1)
	p2 := s2.PointFromLatLng(latlng2)
	p3 := s2.PointFromLatLng(latlng3)
	p4 := s2.PointFromLatLng(latlng4)
	points := []s2.Point{p1, p2, p3, p4}
	loop1 := s2.LoopFromPoints(points)

	//--------------------------
	latlng11 := s2.LatLngFromDegrees(32.78636128858815, -96.79903239011763)
	latlng22 := s2.LatLngFromDegrees(32.78627334662084,-96.79901361465453)
	latlng33 := s2.LatLngFromDegrees(32.78616962008609,-96.79882854223251)
	latlng44 := s2.LatLngFromDegrees(32.786228248142336,-96.79877758026123)
	latlng55 := s2.LatLngFromDegrees(32.78632971968681,-96.79874271154404)
	latlng66 := s2.LatLngFromDegrees(32.78643570095412,-96.79890364408493)

	p11 := s2.PointFromLatLng(latlng11)
	p22 := s2.PointFromLatLng(latlng22)
	p33 := s2.PointFromLatLng(latlng33)
	p44 := s2.PointFromLatLng(latlng44)
	p55 := s2.PointFromLatLng(latlng55)
	p66 := s2.PointFromLatLng(latlng66)
	points1 := []s2.Point{p11, p22, p33, p44, p55, p66}
	loop11 := s2.LoopFromPoints(points1)

	shapeIndex := s2.NewShapeIndex()
	id1 := shapeIndex.Add(loop1)
	id2 := shapeIndex.Add(loop11)

	latlngp := s2.LatLngFromDegrees(32.78641315175891,-96.79930061101913)
	point := s2.PointFromLatLng(latlngp)
	target := s2.NewMaxDistanceToPointTarget(point)

	var opts = s2.NewClosestEdgeQueryOptions().
		MaxResults(2).
		DistanceLimit(s1.ChordAngleFromAngle(s1.Angle(calculate(0.1)))).
		MaxError(s1.ChordAngleFromAngle(s1.Angle(calculate(0.01))))
	query := s2.NewClosestEdgeQuery(shapeIndex, opts)
	result := query.FindEdges(target)
	for _, v := range result {
		if v.ShapeID() == id1 {
			fmt.Println("aaaa")
		}
		if v.ShapeID() == id2 {
			fmt.Println("bbbb")
		}
	}

	containsPointQuery := s2.NewContainsPointQuery(shapeIndex, s2.VertexModelOpen)
	fmt.Println(containsPointQuery.Contains(point))
}

func calculate(distanceKm float64) float64 {
	return distanceKm / 6371
}