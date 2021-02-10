package go_unit_test_bootcamp
//import "fmt"
func FindMissingDrone(droneIds []int) int {
    // TODO: implement this
    var length int = len(droneIds)
    //map that holds the value and count of a number in droneIds
    dict := make(map[int]int)
    //If there is 1 element in the slice, that is our unique ID
    if( length == 1 ){
	    return droneIds[0]
    }
    //Count number of instances of an ID and place into map
    for i := 0; i < length; i++ {
	    dict[droneIds[i]] = dict[droneIds[i]] + 1
    }
    //Iterate through map to check for ID with count of 1
    for value, count := range dict{
	    if(count == 1){
		    return value;
	    }
    }

    return -1
}
