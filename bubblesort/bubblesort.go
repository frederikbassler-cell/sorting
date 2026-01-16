package bubblesort

// BubbleUp geht einmal durch das Array arr und tauscht benachbarte Elemente,
// wenn das linke größer ist als das rechte.
// Gibt true zurück, wenn mindestens ein Tausch stattgefunden hat.
func BubbleUp(arr []int) bool {

	swapped := false

	for i, n := range arr {

		if i+1 < len(arr) {

			if n > arr[i+1] {

				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}

	return swapped
}

// BubbleSort sortiert die übergebene Liste mittels des Bubble-Sort-Algorithmus.
func BubbleSort(arr []int) {

	for d := 0; d <= len(arr); d++ {


		for i, n := range arr {

			if i+1 < len(arr) {

				if n > arr[i+1] {

					arr[i], arr[i+1] = arr[i+1], arr[i]

				}
			}
		}
	}

}
