package go_unit_test_bootcamp

func FindMissingDrone(droneIds []int) int {
	unique := droneIds[0]
	for i, ids := range droneIds{
		dup := false
		for j,neighbor := range droneIds{
			if ids == neighbor{
				if i != j {
					dup = true
					break;
				}
			}
		}		
		if dup == false {
			unique = ids
		}
	}
	return unique
}
