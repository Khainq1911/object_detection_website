import { createContext, useState } from "react";

const filterContext = createContext();

function FilterProvider({ children }) {
  const [openFilter, setOpenFilter] = useState(false);

  const toggleFilter = () => {
    setOpenFilter(!openFilter);
  };
  const closeFilter = () => {
    setOpenFilter(false);
  };

  const value = { openFilter, toggleFilter, closeFilter };
  return (
    <filterContext.Provider value={value}>{children}</filterContext.Provider>
  );
}

export { filterContext, FilterProvider };
