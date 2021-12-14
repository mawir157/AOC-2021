import AdventHelper

import Data.List
import Data.List.Split
import Data.Maybe

import qualified Data.Map as Map

type MapSS = Map.Map String String
type MapSI = Map.Map String Int
type MapCI = Map.Map Char Int

stringToPairs :: MapSI -> String -> MapSI
stringToPairs m [] = m
stringToPairs m [c] = m
stringToPairs m (x:y:xs) = stringToPairs (incr 1 m [x,y]) (y:xs) 

parseRule :: String -> (String, String)
parseRule s = (l, c)
  where [l, c] =  splitOn " -> " s

compExpand :: MapSS -> MapSI -> (String, Int) -> MapSI
compExpand rules cur (s,v) = cur2
  where c = fromJust $ Map.lookup s rules -- B
        cur' = decr s v cur-- remove CH
        cur1 = incr v cur' ([s!!0] ++ c) -- add CB
        cur2 = incr v cur1 (c ++ [s!!1]) -- add BH

compExpRep :: Int -> MapSS -> MapSI -> MapSI
compExpRep 0 _ m = m
compExpRep n rules m = compExpRep (n-1) rules m'
  where m' = foldl (compExpand rules) m $ Map.toList m

score :: [(Char, Int)] -> Int
score f = t - b
  where t = maximum $ map snd f
        b = minimum $ map snd f

freqTable :: MapCI -> [(String, Int)] -> MapCI
freqTable mci [] = mci
freqTable mci ((s, v):xs) = freqTable mci' xs
  where mci' = incr v mci (s!!0)

main = do
  putStrLn "Day 14"
  f <- readFile "../input/input14.txt"
  let p = head $ lines f
  let rs = Map.fromList $ map parseRule $ drop 2 $ lines f

  let p2     = Map.empty :: MapSI
  let counts = Map.empty :: MapCI
  let pMap = stringToPairs p2 p

  let q1 = freqTable counts $ Map.toList $ compExpRep 10 rs pMap
  let r1 = incr 1 q1 (last p) -- we haven't counted the last char yet

  let q2 = freqTable counts $ Map.toList $ compExpRep 40 rs pMap
  let r2 = incr 1 q2 (last p) -- we haven't counted the last char yet

  printSoln 1 $ score $ Map.toList r1
  printSoln 2 $ score $ Map.toList r2
