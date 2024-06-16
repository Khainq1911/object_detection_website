import Header from "~/components/Header";
import Footer from "~/components/Footer";

function Defaultlayout() {
  return (
    <div>
      <Header />
      <div className="content">
        <h1>body</h1>
      </div>
      <Footer />
    </div>
  );
}
export default Defaultlayout;
