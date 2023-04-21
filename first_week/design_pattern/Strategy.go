package main

///google map struct use the route strategy attribute instead of directly implement these
type GoogleMap struct{
	routeStrategy RouteStrategy
}
func (g GoogleMap)buildRoute(start,dest string){
	g.routeStrategy.buildRoute(start,dest)
}

///sepearte strategy from the struct using it
type RouteStrategy interface{
	buildRoute(start,dest string)
}

//different strategy for same operation (find route), can change in runtime
type CarRouteStrategy struct{}
func (s CarRouteStrategy)buildRoute(start,dest string){}

type CyclistRouteStrategy struct{}
func (s CyclistRouteStrategy)buildRoute(start,dest string){}

