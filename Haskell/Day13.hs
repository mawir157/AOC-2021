import AdventHelper

import Data.List
import Data.List.Split

parsePoints :: String -> (Int, Int)
parsePoints s = (read (ps!!0) :: Int, read (ps!!1) :: Int)
  where ps = splitOn "," s

parseFolds :: String -> (Bool, Int)
parseFolds s = (c == 'y', v)
  where c = head $ drop 11 s
        v = read $ drop 13 s :: Int

paperFold :: (Int, Int) -> (Bool, Int) -> (Int, Int)
paperFold (x, y) (b, z)
  | b         = (x, if' (y > z) (2*z - y) y)
  | otherwise = (if' (x > z) (2*z - x) x, y)

buildRow :: Int -> [(Int, Int)] -> String
buildRow n ps = map (\x -> if' ((x,n) `elem` ps) '#' ' ') [ 0..38 ]

buildGrid :: [(Int, Int)] ->  [String]
buildGrid ps = map (\x -> buildRow x ps) [ 0..5 ]

main = do
  putStrLn "Day 13"
  f <- readFile "../input/input13.txt"
  let ps = map parsePoints $ takeWhile (\x -> length x > 0) $ lines f
  let fs = map parseFolds $ drop 1 $ dropWhile (\x -> length x > 0) $ lines f

  printSoln 1 $ length $ nub $ map ( \p -> paperFold p (fs!!0) ) ps
  printSoln 2 $ True
  mapM_ print $ buildGrid $ map ( \p -> foldl paperFold p fs ) ps
