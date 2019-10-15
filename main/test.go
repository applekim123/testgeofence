package main

import (
	"fmt"
	"math"

	"code.uber.internal/growth/jumpops/entity"
	"github.com/golang/geo/s2"
	"github.com/umahmood/haversine"
)
func createRect(lngs []s2.LatLng) *s2.Loop {
	points := make([]s2.Point, len(lngs))
	for _, lng := range lngs {
		points = append(points, s2.PointFromLatLng(lng))
	}
	return s2.LoopFromPoints(points)
}

func getNearPoints(latitude float64, longitude float64, distanceKm float64) []entity.Location {







	dlng := 2 * math.Asin(math.Sin(distanceKm/ (2 * 6371))/math.Cos(latitude * math.Pi / 180))
	dlng = dlng * 180 / math.Pi
	dlat := distanceKm/6371;
	dlat = dlat*180/math.Pi

	result := []entity.Location{}
	result = append(result, entity.Location{latitude + dlat, longitude - dlng})
	result = append(result, entity.Location{latitude - dlat, longitude - dlng})
	result = append(result, entity.Location{latitude - dlat, longitude + dlng})
	result = append(result, entity.Location{latitude + dlat, longitude + dlng})

	return result
}

func main() {
	//vehiclePoint := s2.PointFromLatLng(s2.LatLngFromDegrees(32.78154, -96.796251))

	hh :=haversine.Coord{
		32.788347851915056, -96.79703950881957}
	vv :=haversine.Coord{32.78815844256921, -96.79699659347534}
_, k := haversine.Distance(hh, vv)
fmt.Println(k)


	nearPoints := getNearPoints(32.78154, -96.796251,0.03)
	lngs := []s2.LatLng{}
	for _, nearPoint := range nearPoints {
		lngs = append(lngs, s2.LatLngFromDegrees(nearPoint.Latitude, nearPoint.Longitude))
	}
	loop := createRect(lngs)
	//latlng1 := s2.LatLngFromDegrees(
	//
	//	32.7865468, -96.8010032)
	//latlng2 := s2.LatLngFromDegrees(
	//
	//	32.7858536,-96.8001985)
	//latlng3 := s2.LatLngFromDegrees(
	//	32.7862957,-96.7996733)
	//latlng4 := s2.LatLngFromDegrees(
	//	32.786547,-96.7999687)
	//latlng5 := s2.LatLngFromDegrees(
	//	32.7863607,-96.8001975)
	//latlng6 := s2.LatLngFromDegrees(
	//	32.7867914,-96.8007051)
	latlng1 := s2.LatLngFromDegrees(

		32.78606589343039, -96.79813385009766)
	latlng2 := s2.LatLngFromDegrees(

		32.78536235292114,-96.79723262786865)
	latlng3 := s2.LatLngFromDegrees(
		32.78608393285746,-96.79633140563965)
	latlng4 := s2.LatLngFromDegrees(
		32.786372563193055,-96.79678201675415)
	latlng5 := s2.LatLngFromDegrees(
		32.786047853999655,-96.79723262786865)
	latlng6 := s2.LatLngFromDegrees(
		32.78644472063062,-96.79776906967163)
	p1 := s2.PointFromLatLng(latlng1)
	p2 := s2.PointFromLatLng(latlng2)
	p3 := s2.PointFromLatLng(latlng3)
	p4 := s2.PointFromLatLng(latlng4)
	p5 := s2.PointFromLatLng(latlng5)
	p6 := s2.PointFromLatLng(latlng6)
	points := []s2.Point{p1, p2, p3, p4,p5,p6}
	loop1 := s2.LoopFromPoints(points)



	latlng11 := s2.LatLngFromDegrees(

		32.781809796481774, -96.79657190317037)
	latlng22 := s2.LatLngFromDegrees(

		32.781270203518226,-96.79657190317037)
	latlng33 := s2.LatLngFromDegrees(
		32.781270203518226,-96.79593009682962)
	latlng44 := s2.LatLngFromDegrees(
		32.781809796481774,-96.79593009682962)
	p11 := s2.PointFromLatLng(latlng11)
	p22 := s2.PointFromLatLng(latlng22)
	p33 := s2.PointFromLatLng(latlng33)
	p44 := s2.PointFromLatLng(latlng44)
	points2 := []s2.Point{p11, p22, p33, p44}
	loop2 := s2.LoopFromPoints(points2)

	fmt.Println(loop1.Intersects(loop))
	fmt.Println(loop1.Intersects(loop2))
//	center := haversine.Coord{
//		32.79327460807464, -96.801618039608}
//	leftTop := haversine.Coord{32.79311001147481, -96.80200427770615}
//	_, km := haversine.Distance(center, leftTop)
//	fmt.Println(km)
//	////latlng1 := s2.LatLngFromDegrees(30.250721830490946, -97.74930953979492)
//	////latlng2 := s2.LatLngFromDegrees(30.250652321355673, -97.74909228086472)
//	////latlng3 := s2.LatLngFromDegrees(30.251213026978498, -97.74894207715988)
//	////latlng4 := s2.LatLngFromDegrees(30.251321923980406, -97.74931222200392)
//	////latlng5 := s2.LatLngFromDegrees(30.251442405629117, -97.74809718132019)
//	////latlng1 := s2.LatLngFromDegrees(30.25161386002829, -97.74865508079529)
//	////latlng2 := s2.LatLngFromDegrees(30.251555936278244, -97.7485129237175)
//	////latlng3 := s2.LatLngFromDegrees(30.25167178374416, -97.74840563535689)
//	////latlng4 := s2.LatLngFromDegrees(30.25161386002829, -97.74865508079529)
//	////latlng1 := s2.LatLngFromDegrees(30.251099495933097, -97.74888038635255)
//	////latlng2 := s2.LatLngFromDegrees(30.251041571879757, -97.74873822927476)
//	////latlng3 := s2.LatLngFromDegrees(30.251157419952243, -97.74863094091415)
//	//latlng1 := s2.LatLngFromDegrees(30.250747317161558, -97.74853974580765)
//	//latlng2 := s2.LatLngFromDegrees(30.250557325457933, -97.74848878383636)
//	//latlng3 := s2.LatLngFromDegrees(30.250529521763184, -97.74823933839798)
//	//latlng4 := s2.LatLngFromDegrees(30.250694026842734, -97.7480837702751)
//	//latlng5 := s2.LatLngFromDegrees(30.250842312875587, -97.74833053350449)
//	//p1 := s2.PointFromLatLng(latlng1)
//	//p2 := s2.PointFromLatLng(latlng2)
//	//p3 := s2.PointFromLatLng(latlng3)
//	//p4 := s2.PointFromLatLng(latlng4)
//	//p5 := s2.PointFromLatLng(latlng5)
//	//points := []s2.Point{p1, p2, p3, p4, p5}
//	//loop1 := s2.LoopFromPoints(points)
//	//
//	//cellunion := s2.CellUnion{loop1.CellUnionBound()[0], loop1.CellUnionBound()[1], loop1.CellUnionBound()[2], loop1.CellUnionBound()[3]}
//	//var cellunionpointer *s2.CellUnion
//	//cellunionpointer = &cellunion
//	//
//	//latlngp := s2.LatLngFromDegrees( 30.250608298877907, -97.74837613105774)
//	//cellid := s2.CellIDFromLatLng(latlngp)
//	//paddedcellfromcellid := s2.PaddedCellFromCellID(cellid, 22.0)
//	//p := s2.PointFromLatLng(latlngp)
//	//fmt.Println(paddedcellfromcellid.CellID().ToToken())
//	//fmt.Println(cellid.ToToken())
//	//fmt.Println(loop1.CellUnionBound()[3].ToToken())
//	//fmt.Println(cellunionpointer.ContainsPoint(p))
//	//
//	//cellunionpointer.ExpandAtLevel(20)
//	//
//	////fmt.Println(loop1.CellUnionBound()[0].Level())
//	////fmt.Println(loop1.CellUnionBound()[1].Level())
//	////fmt.Println(loop1.CellUnionBound()[2].Level())
//	////fmt.Println(loop1.CellUnionBound()[3].Level())
//	////loops := []*s2.Loop{loop1}
//	////polygon1 := s2.PolygonFromLoops(loops)
//	//
//	////latlngp := s2.LatLngFromDegrees( 30.25077048685638, -97.74812936782837)
//	////p := s2.PointFromLatLng(latlngp)
//	////fmt.Println(cellunionpointer.ContainsPoint(p))
//	////for _, aa := range cellunion {
//	//////	fmt.Println(aa.ToToken())
//	////}
//	//cell1 := s2.CellFromPoint(p)
//	//
//	//
//	//latlng11 := s2.LatLngFromDegrees( 30.25123619656347, -97.74926126003265)
//	//latlng22 := s2.LatLngFromDegrees(30.250909504910545, -97.74984866380692)
//	//latlng33 := s2.LatLngFromDegrees(30.25036965036558, -97.74951338768004)
//	//latlng44 := s2.LatLngFromDegrees(30.250429891792336, -97.74860948324203)
//	//latlng55 := s2.LatLngFromDegrees(30.25101145135855, -97.74851560592651)
//	//p11 := s2.PointFromLatLng(latlng11)
//	//p22 := s2.PointFromLatLng(latlng22)
//	//p33 := s2.PointFromLatLng(latlng33)
//	//p44 := s2.PointFromLatLng(latlng44)
//	//p55 := s2.PointFromLatLng(latlng55)
//	//points1 := []s2.Point{p11, p22, p33, p44, p55}
//	//loop11 := s2.LoopFromPoints(points1)
//	////loops := []*s2.Loop{loop1}
//	////polygon1 := s2.PolygonFromLoops(loops)
//	//
//	////latlngp := s2.LatLngFromDegrees( 30.251046205805313, -97.74911642074585)
//	////p := s2.PointFromLatLng(latlngp)
//	//
//	//fmt.Println(loop11.Intersects(loop1))
//	//fmt.Println(loop1.IntersectsCell(cell1))
//	//latlng1 := s2.LatLngFromDegrees(30.250721830490946, -97.74930953979492)
//	//latlng2 := s2.LatLngFromDegrees(30.250652321355673, -97.74909228086472)
//	//latlng3 := s2.LatLngFromDegrees(30.251213026978498, -97.74894207715988)
//	//latlng4 := s2.LatLngFromDegrees(30.251321923980406, -97.74931222200392)
//	//latlng5 := s2.LatLngFromDegrees(30.251442405629117, -97.74809718132019)
//	//latlng1 := s2.LatLngFromDegrees(30.25161386002829, -97.74865508079529)
//	//latlng2 := s2.LatLngFromDegrees(30.251555936278244, -97.7485129237175)
//	//latlng3 := s2.LatLngFromDegrees(30.25167178374416, -97.74840563535689)
//	//latlng4 := s2.LatLngFromDegrees(30.25161386002829, -97.74865508079529)
//	//latlng1 := s2.LatLngFromDegrees(30.251099495933097, -97.74888038635255)
//	//latlng2 := s2.LatLngFromDegrees(30.251041571879757, -97.74873822927476)
//	//latlng3 := s2.LatLngFromDegrees(30.251157419952243, -97.74863094091415)
//
//	latlng1 := s2.LatLngFromDegrees(
//		32.7891862,-96.8075261)
//	latlng2 := s2.LatLngFromDegrees(
//		32.7895769,-96.8070012)
//	latlng3 := s2.LatLngFromDegrees(
//		32.7891643,-96.806558)
//	latlng4 := s2.LatLngFromDegrees(
//		32.7887742,-96.8070826)
//	//latlng5 := s2.LatLngFromDegrees(30.250842312875587, -97.74833053350449)
//	p1 := s2.PointFromLatLng(latlng1)
//	p2 := s2.PointFromLatLng(latlng2)
//	p3 := s2.PointFromLatLng(latlng3)
//	p4 := s2.PointFromLatLng(latlng4)
//	//p5 := s2.PointFromLatLng(latlng5)
//
//	c1 := s2.CellIDFromLatLng(latlng1)
//	c2 := s2.CellIDFromLatLng(latlng2)
//	c3 := s2.CellIDFromLatLng(latlng3)
//	c4 := s2.CellIDFromLatLng(latlng4)
//	//c5 := s2.CellIDFromLatLng(latlng5)
//
//	var aaa []s2.CellID
//	aaa = append(aaa, c1)
//	aaa = append(aaa, c2)
//	aaa = append(aaa, c3)
//	aaa = append(aaa, c4)
////	aaa = append(aaa, c5)
//	//var bb s2.CellUnion
//	//bb = aaa
//
//	points := []s2.Point{p1, p2, p3, p4}
//	loop1 := s2.LoopFromPoints(points)
//
//
//
//
//fmt.Println(" ")
//	latlngp := s2.LatLngFromDegrees( 32.7891776404855, -96.8069663643837)
//	pp := s2.PointFromLatLng(latlngp)
//
//fmt.Println(loop1.ContainsPoint(pp))
//	//latlng111 := s2.LatLngFromDegrees(32.797813, -96.807201)
//	//
//	//latlng123 := s2.LatLngFromDegrees(30.250951210288424, -97.74736493825912)
//	//cell123 := s2.CellFromLatLng(latlng123)
//	//
//	//cell123.CapBound()
//	//
//	//
//	//
//	//rect := s2.RectFromCenterSize(latlngp, latlng111)
//	//fmt.Println("aaaaaaaaaaa")
//	//fmt.Println(rect.Lat)
//	//fmt.Println(rect.Lng)
//	//fmt.Println(rect.ContainsLatLng(latlng123))
//	//fmt.Println("aaaaaaaaaaa")
//	//point1 :=s2.PointFromLatLng(latlngp)
//	//cellids := point1.CellUnionBound()
//	//var a []s2.CellID
//	//for _, i := range cellids {
//	//	a = append(a, i)
//	//}
//	//var b s2.CellUnion
//	//b = a
//	//b.ExpandAtLevel(23)
//	//for _, i1 := range b {
//	//	fmt.Print(i1.ToToken())
//	//	fmt.Print(",")
//	//}
//	//cellid := s2.CellIDFromLatLng(latlngp)
//	//cell := s2.CellFromCellID(cellid)
//	//
//	//fmt.Println(loop1.IntersectsCell(cell))
//	//
//	//latlng11 := s2.LatLngFromDegrees( 30.25123619656347, -97.74926126003265)
//	//latlng22 := s2.LatLngFromDegrees(30.250909504910545, -97.74984866380692)
//	//latlng33 := s2.LatLngFromDegrees(30.25036965036558, -97.74951338768004)
//	//latlng44 := s2.LatLngFromDegrees(30.250429891792336, -97.74860948324203)
//	//latlng55 := s2.LatLngFromDegrees(30.25101145135855, -97.74851560592651)
//	//p11 := s2.PointFromLatLng(latlng11)
//	//p22 := s2.PointFromLatLng(latlng22)
//	//p33 := s2.PointFromLatLng(latlng33)
//	//p44 := s2.PointFromLatLng(latlng44)
//	//p55 := s2.PointFromLatLng(latlng55)
//	//points1 := []s2.Point{p11, p22, p33, p44, p55}
//	//loop11 := s2.LoopFromPoints(points1)
//	//
//	//
//	//fmt.Println(loop11.Intersects(loop1))


}

