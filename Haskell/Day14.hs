import AdventHelper

import Data.List
import Data.List.Split
import Data.Maybe

import qualified Data.Map as Map

type MapSS = Map.Map String String
type MapSI = Map.Map String Int
type MapCI = Map.Map Char Int

incr :: (Ord k) => k -> Int -> Map.Map k Int -> Map.Map k Int
incr k i m
  | Map.member k m = Map.adjust (\v -> v + i) k m
  | otherwise      = Map.insert k i m

decr :: (Ord k) => k -> Int -> Map.Map k Int -> Map.Map k Int
decr k d m
  | fromJust (Map.lookup k m) > d = Map.adjust (\v -> v - d) k m
  | otherwise                     = Map.delete k m

stringToPairs :: MapSI -> String -> MapSI
stringToPairs m [] = m
stringToPairs m [c] = m
stringToPairs m (x:y:xs) = stringToPairs (incr [x,y] 1 m) (y:xs) 

parseRule :: String -> (String, String)
parseRule s = (l, c)
  where [l, c] =  splitOn " -> " s

expand :: Map.Map String String -> String -> String
expand _ [] = []
expand _ [c] = [c]
expand m (x:y:xs)
  | v == Nothing = [x] ++ (expand m (y:xs))
  | otherwise    = ([x] ++ (fromJust v)) ++ (expand m (y:xs))
  where v = Map.lookup [x,y] m

expandRep :: Int -> Map.Map String String -> String -> String
expandRep 0 _ s = s
expandRep n m s = expandRep (n-1) m $ expand m s

-- e.g. CH -> B == (CH, B)
compExpand :: MapSS -> MapSI -> (String, Int) -> MapSI
compExpand rules cur (s,v) = cur2
  where c = fromJust $ Map.lookup s rules -- B
        cur' = decr s v cur-- remove CH
        cur1 = incr([s!!0] ++ c) v cur'-- add CB
        cur2 = incr (c ++ [s!!1]) v cur1-- add BH

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
  where mci' = incr (s!!0) v mci

main = do
  putStrLn "Day 14"
  f <- readFile "../input/input14.txt"
  let p = head $ lines f
  let rs = Map.fromList $ map parseRule $ drop 2 $ lines f

  let p2     = Map.empty :: MapSI
  let counts = Map.empty :: MapCI
  let pMap = stringToPairs p2 p

  let q = freqTable counts $ Map.toList $ compExpRep 40 rs pMap
  let q' = incr (last p) 1 q -- we haven't counted the last char yet

  printSoln 1 $ score $ freq $ expandRep 10 rs p
  printSoln 2 $ score $ Map.toList q'
