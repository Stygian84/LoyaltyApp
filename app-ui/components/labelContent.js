const LabelContent = ({ title, children }) => {
  return (
    <div className="mb-4 flex flex-col justify-center items-center">
      {/* vertical align p tag  */}

      <label className="mr-2 mb-2 font-bold ">{title}:</label>
      {children}
    </div>
  );
};
export default LabelContent;
