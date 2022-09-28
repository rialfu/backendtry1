import React, {useState, useEffect} from 'react'

export default function Home() {
    const [data, setData] = useState(null);
    const [error, setError] = useState(null);
    useEffect(() => {
        getData()
    }, []);

    const getData = async () => {
        try {
            // const res = await fetch("https://backendtry1.herokuapp.com/")
            const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL);
            const data = await res.json();
            setData(data);
        }
        catch (error) {
          setError(error);
        }
    }
    return (
        <div>
            <Input onSuccess={getData} />
            {error && <div>Failed to load {error.toString()}</div>}
            {!data ? 
                <div>Loading...</div> : 
                ((data?.data ?? []).length === 0 && <p>data kosong</p>)
            }
            {data?.data ? data.data.map((item, index) => (
        <p key={index}>{item}</p>
      )) :
        <p>data kosong</p>
      }
    
        </div>
    )
}
function Input({onSuccess}) {
    const [data, setData] = useState(null);
    const [error, setError] = useState(null);
    const handleSubmit = async (e) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);
        const body = {
          text: formData.get("data")
        }
    
        try {
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/send`, {
                method: 'POST',
                body: JSON.stringify(body)
            });
            const data = await res.json();
            setData(data.message);
            onSuccess();
        }
        catch (error) {
            setError(error);
        }
    }
    return (
        <div>
            <form onSubmit={handleSubmit}>
            <input name="data" type="text" />
            <button >Submit</button>
            </form>
        </div>
    )
  }
