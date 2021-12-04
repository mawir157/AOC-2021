import AdventHelper

import Data.List.Split

parseNumbers :: String -> [Int]
parseNumbers s = map read $ splitOn "," s :: [Int]

parseBoards :: [String] -> [[Int]]
parseBoards [] = []
parseBoards ss = [(buildBoard (take 5 ss))] ++ (parseBoards (drop 6 ss))

buildBoard :: [String] -> [Int]
buildBoard [] = []
buildBoard (s:ss) = s' ++ buildBoard ss
  where s' = map read $ filter (\t -> t /= "") $ splitOn " " s :: [Int]

callNumber :: Int -> [Int] -> [Int]
callNumber n b = map (\x -> if' (x == n) (-1) x) b

winRows :: [Int] -> Bool
winRows [] = False
winRows xs = ((sum (take 5 xs)) == -5) || (winRows (drop 5 xs))

everyf n [] = []
everyf n as  = head as : everyf n (drop n as)

winCols :: Int -> [Int] -> Bool
winCols 5 _ = False
winCols _ [] = False
winCols n xs = ((sum f) == -5) || (winCols (n+1) (drop 1 xs))
  where f = everyf 5 xs

winning :: [Int] -> Bool
winning b = (winRows b) || (winCols 0 b)

score :: [Int] -> Int
score b = sum $ filter (> 0) b

part1 :: [[Int]] -> [Int] -> Int
part1 bs (n:ns)
  | length win == 1 = n * (score (head win))
  | otherwise       = part1 bs' ns
  where bs' = map (callNumber n) bs
        win = filter winning bs'

part2 :: [[Int]] -> [Int] -> Int
part2 bs (n:ns)
  | length bs' == 1 = part1 bs' ns
  | otherwise       = part2 bs' ns
  where bs' = filter (not.winning) $ map (callNumber n) bs

main = do
  putStrLn "Day 4"
  f <- readFile "../input/input04.txt"
  let ss = lines f

  let calls = parseNumbers (head ss)
  let boards = parseBoards (drop 2 ss)

  printSoln 1 $ part1 boards calls
  printSoln 2 $ part2 boards calls
