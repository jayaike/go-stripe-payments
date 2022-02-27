import React, { useEffect, useState } from 'react';
import { Elements, PaymentElement, useElements, useStripe } from '@stripe/react-stripe-js';
import { loadStripe } from '@stripe/stripe-js'


const stripePromise = loadStripe('Your stripe public key');

function App() {
  const [clientSecret, setClientSecret] = useState(null);

  useEffect(() => {
    fetch('http://localhost:8080/payments/cart/initialise/', {
      method: 'POST',
      body: JSON.stringify({
        itemId: '1234'
      })
    })
    .then(res => res.json())
    .then(({ data }) => setClientSecret(data.clientSecret))
  }, [])

  const options = {
    // passing the client secret obtained in step 2
    clientSecret: clientSecret, // This is returned from out api
    appearance: {
      variables: {
        colorPrimaryText: '#FFF',
        colorText: '#FFF',
        colorBackground: '#0f0f24'
      }
    },
  };

  if(!clientSecret) {
    return <></>
  }

  return (
    <Elements stripe={stripePromise} options={options}>
        <CheckoutForm />
    </Elements>
  );
}


const CheckoutForm = () => {
  const stripe = useStripe();
  const elements = useElements();

  const [errorMessage, setErrorMessage] = useState(null);

  const handleSubmit = async (event) => {
    // We don't want to let default form submission happen here,
    // which would refresh the page.
    event.preventDefault();

    if (!stripe || !elements) {
      // Stripe.js has not yet loaded.
      // Make sure to disable form submission until Stripe.js has loaded.
      return;
    }

    const {error} = await stripe.confirmPayment({
      //`Elements` instance that was used to create the Payment Element
      elements,
      confirmParams: {
        return_url: 'https://my-site.com/order/123/complete',
      },
    });


    if (error) {
      // This point will only be reached if there is an immediate error when
      // confirming the payment. Show error to your customer (for example, payment
      // details incomplete)
      setErrorMessage(error.message);
    } else {
      // Your customer will be redirected to your `return_url`. For some payment
      // methods like iDEAL, your customer will be redirected to an intermediate
      // site first to authorize the payment, then redirected to the `return_url`.
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      style={{
        width: '350px',
        margin: 'auto',
        marginTop: '45px',
        padding: '50px',
        borderRadius: '25px',
        background: '#1A1A40'
      }}>
        <PaymentElement />
        <button
          disabled={!stripe}
          style={{
            marginTop: '15px',
            width: '100%',
            padding: '12.5px 15px',
            background: '#635BFF',
            border: 'none',
            outline: 'none',
            borderRadius: '5px',
            color: '#FFF',
            fontWeight: 'bold',
            opacity: !stripe ? '0.4' : '1'
          }}
        >
          Pay $29.99
        </button>
        {/* Show error message to your customers */}
        {errorMessage && <div>{errorMessage}</div>}
    </form>
  )
};

export default App;
