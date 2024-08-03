import React, { useState } from "react";
import classNames from "classnames/bind";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faChevronDown, faChevronUp } from "@fortawesome/free-solid-svg-icons";
import styles from "./menu.modules.scss";

const cx = classNames.bind(styles);

function Menu({ data, isOpen, onToggle, getValue }) {
  const [text, setText] = useState(null);

  const setValue = (value) => {
    setText(value);
    onToggle();
  };

  return (
    <div className={cx("menu_container")} onClick={onToggle}>
      <div className={cx("menu_label")}>
        <p>{text || "Select"}</p>
        <FontAwesomeIcon
          icon={isOpen ? faChevronUp : faChevronDown}
          className={cx("menu_icon")}
        />
      </div>
      {isOpen && (
        <ul className={cx("menu_dropdown")}>
          {data.map((val, index) => (
            <li
              key={index}
              onClick={() => {
                setValue(val);
                getValue(val);
              }}
            >
              {val}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default Menu;
