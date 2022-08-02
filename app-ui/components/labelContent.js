const LabelContent = ({ title, children }) => {
  return (
    <div className="flex flex-col justify-center items-center">
      {/* vertical align p tag  */}

      <label className="mr-2 font-bold ">{title}:</label>
      {children}
    </div>
  );
};
export default LabelContent;
