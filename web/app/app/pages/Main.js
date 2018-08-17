import Numpad from 'molecules/Numpad';

export default function Main() {
  return (
    <div className="row">
      <div className="8 col">
        <table className="w-100">
          <tr>
            <td>1</td>
            <td>2</td>
            <td>3</td>
          </tr>
          <tr>
            <td>1</td>
            <td>2</td>
            <td>3</td>
          </tr>
          <tr>
            <td>1</td>
            <td>2</td>
            <td>3</td>
          </tr>
        </table>
      </div>
      <div className="4 col">
        <Numpad onClick={pressed} />
      </div>
    </div>
  );
}

function pressed(n) {
  return n + 1;
}
