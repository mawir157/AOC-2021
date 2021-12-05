import AdventHelper

import Data.List.Split
import qualified Data.Map as Map

type Line = ((Int, Int), (Int, Int))
type Space = Map.Map (Int, Int) Int

parseLine :: String -> ((Int, Int), (Int, Int))
parseLine s = ((start!!0, start!!1), (end!!0, end!!1))
  where [lhs, rhs] = splitOn " -> " s
        start = map read $ splitOn "," lhs :: [Int]
        end = map read $ splitOn "," rhs :: [Int]

range :: (Int, Int) -> [Int]
range (x,y) = [x, (x + if' (x < y) 1 (-1)) .. y]

linePoints :: Bool -> Line -> [(Int, Int)]
linePoints b ((x0,y0),(x1,y1))
 | x0 == x1  = [ (x0, vy) | vy <- range (y0, y1)]
 | y0 == y1  = [ (vx, y0) | vx <- range (x0, x1)]
 | b         = [ (vx, vy) | (vx, vy) <- zip (range (x0, x1)) (range (y0, y1))]
 | otherwise = []

increment :: Ord a => Map.Map a Int -> a -> Map.Map a Int
increment m k
  | Map.member k m = Map.adjust (\v -> v + 1) k m
  | otherwise      = Map.insert k 1 m

addLine :: Bool -> Space -> Line -> Space
addLine b m l = foldl (increment) m $ linePoints b l

main = do
  putStrLn "Day 5"
  f <- readFile "../input/input05.txt"
  let ss = lines f
  let ls = map parseLine ss

  let s = Map.empty :: Space
  let r1 = Map.filter (> 1) $ foldl (addLine False) s ls
  let r2 = Map.filter (> 1) $ foldl (addLine True) s ls

  printSoln 1 $ Map.size r1
  printSoln 2 $ Map.size r2
