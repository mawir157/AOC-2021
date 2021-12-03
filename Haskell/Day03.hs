import AdventHelper

bitCount :: [String] -> [Int]
bitCount ss = foldl1 (zipWith (+)) $ map singleCount ss
  where singleCount s = map (\c -> if' (c == '1') 1 0) s

toBin :: [Bool] -> Int
toBin b = foldl toBin' 0 b
  where toBin' x b = 2 * x + (if' b 1 0)

stringToBin :: String -> Int
stringToBin ss = foldl stringToBin' 0 ss
  where stringToBin' x c = 2 * x + (if' (c == '1') 1 0)

reduce :: Bool -> Int -> [String] -> String
reduce _ _ [s] = s
reduce b n ss = reduce b (n+1) ss'
  where bc = bitCount ss
        q = (2 * (bc!!n)) >= (length ss)
        t = if' q '1' '0'
        ss' = filter (\w -> (w!!n /= t) == b) ss

main = do
  putStrLn "Day 3"
  f <- readFile "../input/input03.txt"
  let ss = lines f

  let gamma = toBin $ map (\x -> 2*x >= length ss) $ bitCount ss
  let epsilon = toBin $ map (\x -> 2*x < length ss) $ bitCount ss
  
  let o2 = stringToBin $ reduce True 0 ss
  let co2 = stringToBin $ reduce False 0 ss

  printSoln 1 $ (gamma * epsilon)
  printSoln 2 $ (o2 * co2)
