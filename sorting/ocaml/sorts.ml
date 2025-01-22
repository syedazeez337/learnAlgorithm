

(* Function to insert an element into a sorted list *)
let rec insert x sorted_list =
  match sorted_list with
  | [] -> [x]
  | h :: t ->
    if x <= h then x :: sorted_list
    else h :: insert x t

let rec insertion_sort lst =
  match lst with
  | [] -> []
  | h :: t -> insert h (insertion_sort t)