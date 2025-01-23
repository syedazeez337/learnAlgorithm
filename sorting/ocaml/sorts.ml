

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


(* Merge two sorted lists into one sorted list *)
let rec merge left right =
  match (left, right) with
  | [], _ -> right (* If left is empty, return right *)
  | _, [] -> left  (* If right is empty, return left *)
  | h1 :: t1, h2 :: t2 ->
      if h1 <= h2 then h1 :: merge t1 right (* Append smaller element *)
      else h2 :: merge left t2

(* Merge sort implementation *)
let rec merge_sort lst =
  match lst with
  | [] | [_] -> lst (* Base case: empty or single-element list is sorted *)
  | _ ->
      let mid = List.length lst / 2 in
      let left = List.filteri (fun i _ -> i < mid) lst in
      let right = List.filteri (fun i _ -> i >= mid) lst in
      merge (merge_sort left) (merge_sort right)
