import styles from "./Footer.module.scss"
import classNames from "classnames/bind"
const cx =classNames.bind(styles)
function Footer(){
    return (
        <div className={cx("footer_container")}>
            <h1>IVSR lab</h1>
        </div>
    )
}
export default Footer