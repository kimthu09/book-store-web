export const LoadingSpinner = () => {
  return (
    <div className="fixed top-0 left-0 w-full h-full  flex justify-center bg-[rgba(212,204,204,0.31)] items-center z-[999]">
      <div className="p-6 rounded-lg bg-[rgba(104,104,111,0.3)]  shadow-md">
        <div className="border-4 border-white border-t-primary rounded-full w-10 h-10 animate-spin"></div>
      </div>
    </div>
  );
};
