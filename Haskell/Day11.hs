import AdventHelper

import qualified Data.Map as Map

type Octos = Map.Map (Int,Int) (Int, Bool)

dirs :: [(Int,Int)]
dirs = [(-1,-1), (-1,0), (-1,1), (0,-1), (0,1), (1,-1), (1,0), (1,1)]

parseInput :: Int -> [String] -> Octos
parseInput _ [] = Map.empty :: Octos
parseInput n (s:ss) = Map.union l (parseInput (n+1) ss)
  where l = Map.fromList $ parseLine n 0 s

parseLine :: Int -> Int -> String -> [((Int, Int), (Int, Bool))]
parseLine _ _ [] = []
parseLine r c (s:ss) = [((r,c), (v, False))]  ++ (parseLine r (c+1) ss)
  where v = read [s] :: Int

octoTick :: (Octos, Int) -> (Octos, Int)
octoTick (os, n) = (os', n')
  where ps = foldl incHelper os $ Map.keys os -- increment all octopus
        (qs, n') = flashRep (ps, n) -- do all flashes
        os' = foldl resetHelper qs $ Map.keys qs -- reset all to unflashed

flash :: (Octos, Int) -> (Octos, Int)
flash (os, n) = (os', n + length fs)
  where fs = Map.keys $ Map.filter (flashNow) os -- which octopus flash
        nbrs = concat $ map getNbrs fs -- the octopus hit by flashes
        ps = foldl incHelper os nbrs -- increment the hit octopus
        os' = foldl flashedHelper ps fs -- flag the flashed octopuss
        flashNow (v,b) = (v > 9) && (not b)
        getNbrs (x,y) = map (\(dx, dy) -> (x+dx, y+dy)) dirs

flashRep :: (Octos, Int) -> (Octos, Int)
flashRep (os, n)
  | n == n'   = (os, n)
  | otherwise = flashRep $ (os', n')
  where (os', n') = flash (os, n)

incHelper :: Octos -> (Int, Int) -> Octos
incHelper os p = Map.adjust (\(v,b) -> (v+1,b)) p os

flashedHelper :: Octos -> (Int, Int) -> Octos
flashedHelper os p = Map.adjust(\(v,b) -> (v, True)) p os

resetHelper :: Octos -> (Int, Int) -> Octos
resetHelper os p = Map.adjust(\(v,b) -> (if' (v > 9) 0 v, False)) p os

octoLife :: Int -> (Octos, Int) ->  (Octos, Int)
octoLife 0 os = os
octoLife n os = octoLife (n-1) $ octoTick os

octoLifeSim :: (Octos, Int) -> (Octos, Int)
octoLifeSim (os, n)
  | flag      = (os, n)
  | otherwise = octoLifeSim (os', n+1)
  where ps = foldl incHelper os $ Map.keys os -- increment all octopus
        (qs, n') = flashRep (ps, n) -- do all flashes
        flag = and $ map snd $ Map.elems qs -- everything flashed this tick
        os' = foldl resetHelper qs $ Map.keys qs 

main = do
  putStrLn "Day 11"
  f <- readFile "../input/input11.txt"
  let ss = lines f
  let os = parseInput 0 ss

  printSoln 1 $ snd $ octoLife 100 (os, 0)
  printSoln 2 $ snd $ octoLifeSim (os, 1)
