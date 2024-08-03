import classNames from "classnames/bind";
import styles from "./Header.module.scss";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faBars,
  faChartPie,
  faHouse,
  faMagnifyingGlass,
} from "@fortawesome/free-solid-svg-icons";
import { useContext, useState } from "react";
import { filterContext } from "~/hooks/useContext";

const url = "http://localhost:1911/";
const cx = classNames.bind(styles);

function Header() {
  const context = useContext(filterContext);
  const [activeMenu, setActiveMenu] = useState(false);

  const handleOpenMenu = () => {
    setActiveMenu(!activeMenu);
  };

  return (
    <div className={cx("header_container")}>
      <nav className={cx("nav_container")}>
        <ul>
          <li>
            <a href={url}>
              <FontAwesomeIcon icon={faHouse} className={cx("icon")} />
              Home
            </a>
          </li>
          <li>
            <a href={url + `chart`}>
              <FontAwesomeIcon icon={faChartPie} className={cx("icon")} />
              Chart
            </a>
          </li>
          <li>
            <div className={cx("filter_btn")} onClick={context.toggleFilter}>
              <FontAwesomeIcon
                icon={faMagnifyingGlass}
                className={cx("icon")}
              />
              Filter
            </div>
          </li>
        </ul>
      </nav>
      <div className={cx("header_menu_container")} onClick={handleOpenMenu}>
        <FontAwesomeIcon icon={faBars} />
        {activeMenu && (
          <div className={cx("menu")}>
            <p>Account</p>
            <p>Logout</p>
          </div>
        )}
      </div>
    </div>
  );
}
export default Header;
