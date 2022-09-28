import React, {useState, useEffect} from 'react'
interface getData{
    data?: Array<String>
}

export default function Home() {
    const [data, setData] = useState<getData>(null);
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
            setError(null)
        }
        catch (error) {
          setError(error);
        }
    }
    return (
        <div>
            <h1 className="text-center font-bold text-blue text-3xl">Welcome to Record</h1>
            <div className="p-3 border w-1/2 m-auto" >
            <Input onSuccess={getData} />
            </div>
            <div className="pl-10">
            {error ?
                <div>Failed to load {error.toString()}</div>
                :
                !data ?
                    <div>Loading ....</div>
                    : (
                        (data?.data ?? []).length === 0 ? <p>Data kosong</p>
                        : (data?.data?? [ ]).map((data, i)=><p key={i}>{i+1}. {data}</p>)
                    )
            }
            </div>
            
            

            {/* {error && <div>Failed to load {error.toString()}</div>}
            {!data ? 
                <div>Loading...</div> : 
                ((data?.data ?? []).length === 0 && <p>data kosong</p>)
            }
            {data?.data ? data.data.map((item, index) => (
                <p key={index}>{item}</p>
            )) :
                <></>
            }

            <h1 className="text-3xl font-bold underline">
                Hello world!
            </h1> */}


        </div>
    )
}
function Input({onSuccess}) {
    const [val, setVal] = useState('')
    const [data, setData] = useState(null);
    const [error, setError] = useState(null);
    const handleSubmit = async (e) => {
        e.preventDefault()
        console.log(data)
        // e.preventDefault();
        // const formData = new FormData(e.currentTarget);
        if(val === ""){
            alert("data harus dimasukkan")
            return 
        }
        const body = {
          text: val
        }
    
        try {
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/send`, {
                method: 'POST',
                body: JSON.stringify(body)
            });
            const data = await res.json();
            setData(data.message);
            setVal('')
            onSuccess();
        }
        catch (error) {
            setError(error);
        }
    }
    return (
        <div>
            <form onSubmit={handleSubmit} >
            <div className="flex">
            <input type="text" id="default-input" placeholder="input data" className="bg-gray-50 border border-gray-300 
                text-gray-900 text-sm rounded-lg focus:ring-blue-500 w-full
            focus:border-blue-500 block p-2.5 dark:bg-gray-700 dark:border-gray-600 
            dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            onChange={(e)=>setVal(e.target.value)} value={val}
            />
            <button className="bg-blue-400 text-white  rounded-lg ml-1 p-2.5">Submit</button>
            </div>
            
            </form>
        </div>
    )
  }
