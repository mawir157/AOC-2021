import AdventHelper

import Data.Maybe
import Data.List

import qualified Data.Map as Map

type Ocean = Map.Map (Int, Int) Int

parseInput :: Int -> [String] -> Ocean
parseInput _ [] = Map.empty :: Map.Map (Int,Int) Int
parseInput n (s:ss) = Map.union l (parseInput (n+1) ss)
  where l = Map.fromList $ parseLine n 0 s

parseLine :: Int -> Int -> String -> [((Int, Int), Int)]
parseLine _ _ [] = []
parseLine r c (s:ss) = [((r,c), v)]  ++ (parseLine r (c+1) ss)
  where v = read [s] :: Int

thisIsALow :: Ocean -> (Int, Int) -> Bool
thisIsALow os (i,j) = and [u, d, r, l]
  where v = fromJust $ Map.lookup (i,j) os
        u = v < (Map.findWithDefault 9 (i-1, j) os)
        d = v < (Map.findWithDefault 9 (i+1, j) os)
        l = v < (Map.findWithDefault 9 (i, j-1) os)
        r = v < (Map.findWithDefault 9 (i, j+1) os)

risk :: Ocean -> (Int, Int) -> Int
risk os p = (1 + (fromJust $ Map.lookup p os))

bSize :: Ocean -> (Int, Int) -> (Ocean, Int)
bSize os (i,j)
  | not b     = (os, 0)
  | v == 9    = (os, 0)
  | otherwise = (osr, 1 + nu + nd + nl + nr)
  where b = Map.member (i,j) os
        v = fromJust $ Map.lookup (i,j) os
        os' = Map.insert (i,j) 9 os
        (osu, nu) = (bSize os' (i+1,j))
        (osd, nd) = (bSize osu (i-1,j))
        (osl, nl) = (bSize osd (i,j-1))
        (osr, nr) = (bSize osl (i,j+1))

main = do
  putStrLn "Day 9"
  f <- readFile "../input/input09.txt"
  let ss = lines f
  let fl = parseInput 0 ss
  let t = filter (\p -> thisIsALow fl p ) [(i,j) | i <- [0,1..99], j <- [0,1..99]]

  printSoln 1 $ sum $ map (risk fl) t
  printSoln 2 $ product $ take 3 $ reverse $ sort $ map (snd . bSize fl) t
