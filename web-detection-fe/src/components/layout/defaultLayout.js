import Header from "~/components/Header";
import Footer from "~/components/Footer";

function Defaultlayout({ children }) {
  return (
    <div>
      <Header />
      <div className="content">{children}</div>
      <Footer />
    </div>
  );
}
export default Defaultlayout;
