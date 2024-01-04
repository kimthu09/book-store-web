import React from "react";
import { Button } from "./ui/button";
import {
  LuChevronLeft,
  LuChevronRight,
  LuChevronsLeft,
  LuChevronsRight,
} from "react-icons/lu";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select";

export interface PagingProps {
  page: string;
  totalPage: number;
  onNavigateBack: () => void;
  onNavigateNext: () => void;
  onNavigateFirst: () => void;
  onNavigateLast: () => void;
  onPageSelect: (selectedPage: string) => void;
}

const Paging = ({
  page,
  totalPage,
  onNavigateBack,
  onNavigateNext,
  onNavigateFirst,
  onNavigateLast,
  onPageSelect,
}: PagingProps) => {
  const pageArray = Array.from({ length: totalPage }, (value, key) => key + 1);
  return (
    <div className="flex gap-2">
      <Button
        variant="outline"
        size="icon"
        onClick={onNavigateFirst}
        disabled={Number(page) <= 1}
      >
        <LuChevronsLeft className="h-4 w-4" />
      </Button>
      <Button
        variant="outline"
        size="icon"
        onClick={onNavigateBack}
        disabled={Number(page) <= 1}
      >
        <LuChevronLeft className="h-4 w-4" />
      </Button>
      <Select
        value={page}
        onValueChange={(value) => {
          if (value != page) onPageSelect(value);
        }}
      >
        <SelectTrigger className="w-[100px]">
          <SelectValue />
        </SelectTrigger>
        <SelectContent className="min-w-0">
          <SelectGroup>
            {pageArray.map((item) => {
              return (
                <SelectItem key={item} value={`${item}`}>
                  {`Trang ${item}`}
                </SelectItem>
              );
            })}
          </SelectGroup>
        </SelectContent>
      </Select>
      <Button
        variant="outline"
        size="icon"
        onClick={onNavigateNext}
        disabled={Number(page) >= totalPage}
      >
        <LuChevronRight className="h-4 w-4" />
      </Button>
      <Button
        variant="outline"
        size="icon"
        onClick={onNavigateLast}
        disabled={Number(page) >= totalPage}
      >
        <LuChevronsRight className="h-4 w-4" />
      </Button>
    </div>
  );
};

export default Paging;
