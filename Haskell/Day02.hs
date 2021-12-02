import AdventHelper

fst' (x, _, _) = x
snd' (_, x, _) = x

parseRule :: String -> (String, Int)
parseRule s = (takeWhile(/= ' ') s, i)
  where i = read $ drop 1 $ dropWhile(/= ' ') s :: Int

apply :: (Int, Int) -> (String, Int) -> (Int, Int)
apply (h, v) (d, val)
  | d == "forward" = (h + val, v)
  | d == "up"      = (h, v - val)
  | d == "down"    = (h, v + val)

applyWithAim :: (Int, Int, Int) -> (String, Int) -> (Int, Int, Int)
applyWithAim (h, v, a) (d, val)
  | d == "forward" = (h + val, v + (a * val), a)
  | d == "up"      = (h, v, a - val)
  | d == "down"    = (h, v, a + val)

main = do
  putStrLn "Day 2"
  f <- readFile "../input/input02.txt"
  let is = map parseRule $ lines f

  let p1 = foldl apply (0,0) is 
  let p2 = foldl applyWithAim (0,0,0) is 

  printSoln 1 $ (fst p1) * (snd p1)
  printSoln 2 $ (fst' p2) * (snd' p2)
