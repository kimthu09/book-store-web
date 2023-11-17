import { useState, useEffect, useRef, ChangeEvent } from "react";
import { Input } from "./ui/input";
import { Label } from "./ui/label";

const RangeSlider = ({
  initialMin,
  initialMax,
  min,
  max,
  step,
  priceCap,
}: {
  initialMin: number;
  initialMax: number;
  min: number;
  max: number;
  step: number;
  priceCap: number;
}) => {
  const progressRef = useRef<HTMLInputElement>(null);
  const [minValue, setMinValue] = useState(initialMin);
  const [maxValue, setMaxValue] = useState(initialMax);
  const handleMin = (e: ChangeEvent<HTMLInputElement>) => {
    if (maxValue - minValue >= priceCap && maxValue <= max) {
      if (parseInt(e.target.value) > maxValue) {
      } else {
        setMinValue(parseInt(e.target.value));
      }
    } else {
      if (parseInt(e.target.value) < minValue) {
        setMinValue(parseInt(e.target.value));
      }
    }
  };
  const handleMax = (e: ChangeEvent<HTMLInputElement>) => {
    if (maxValue - minValue >= priceCap && maxValue <= max) {
      if (parseInt(e.target.value) < minValue) {
      } else {
        setMaxValue(parseInt(e.target.value));
      }
    } else {
      if (parseInt(e.target.value) > maxValue) {
        setMaxValue(parseInt(e.target.value));
      }
    }
  };
  useEffect(() => {
    if (progressRef.current != null) {
      progressRef.current.style.left = (minValue / max) * step + "%";
      progressRef.current.style.right = step - (maxValue / max) * step + "%";
    }
  }, [minValue, maxValue, max, step]);
  return (
    <div className="grid place-items-center">
      <div className="flex flex-col">
        <div className="flex justify-between items-center mb-6 ">
          <div className="rounded-md mr-24">
            <Label>
              Giá nhỏ nhất <br />
            </Label>
            <Input
              onChange={(e) => setMinValue(parseInt(e.target.value))}
              type="number"
              min={min}
              step={step}
              max={maxValue}
              value={minValue}
              className="w-24 rounded-md border border-gray-400"
            />
          </div>
          <div className=" ">
            <Label>Giá lớn nhất</Label>
            <Input
              onChange={(e) => setMaxValue(parseInt(e.target.value))}
              type="number"
              value={maxValue}
              min={minValue}
              step={step}
              max={max}
              className="w-24 rounded-md border border-blue-400"
            />
          </div>
        </div>

        <div className="mb-4">
          <div className="slider relative h-1 rounded-md bg-blue-300">
            <div
              className="progress absolute h-1 bg-blue-300 rounded "
              ref={progressRef}
            ></div>
          </div>

          <div className="range-input relative  ">
            <input
              onChange={handleMin}
              type="range"
              min={min}
              step={step}
              max={max}
              value={minValue}
              className="range-min absolute w-full  -top-1  h-1   bg-transparent  appearance-none pointer-events-none"
            />

            <input
              onChange={handleMax}
              type="range"
              min={min}
              step={step}
              max={max}
              value={maxValue}
              className="range-max absolute w-full  -top-1 h-1  bg-transparent appearance-none  pointer-events-none"
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default RangeSlider;
