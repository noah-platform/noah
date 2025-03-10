import type { Route } from './+types/index';
import slide_1 from '~/assets/home/slide_1.svg';
import slide_2 from '~/assets/home/slide_2.svg';
import slide_3 from '~/assets/home/slide_3.svg';
import slide_4 from '~/assets/home/slide_4.svg';
import std_1 from '~/assets/home/std_1.svg';
import std_2 from '~/assets/home/std_2.svg';
import std_3 from '~/assets/home/std_3.svg';
import teacher_1 from '~/assets/home/teacher_1.svg';
import teacher_2 from '~/assets/home/teacher_2.svg';
import teacher_3 from '~/assets/home/teacher_3.svg';
import listening from '~/assets/home/listening.svg';
import writing from '~/assets/home/writing.svg';
import reading from '~/assets/home/reading.svg';
import speaking from '~/assets/home/speaking.svg';
import click from '~/assets/home/click.svg';
import dashboard from '~/assets/home/dashboard.svg';
import { useState, useEffect } from 'react';
import { motion } from 'framer-motion';

export function meta({}: Route.MetaArgs) {
  return [{ title: 'NOAH ENGLISH' }];
}

export default function Home() {
  return (
    <>
      <SlideBar />
      <SampleQuestion />
      <ReviewPage />
      <Personalization />
      <InsightAnalytic />
      <PricingPage />
    </>
  );
}

function slide1() {
  return (
    <div className="flex-1 flex flex-col justify-between items-center text-center">
      {/* Title and Description */}
      <div className="flex-grow flex flex-col justify-center items-center">
        <h2 className="text-5xl font-bold text-white leading-20">
          Are you ready for<br className="mb10"></br> the{' '}
          <span className="text-5xl font-bold text-[#EF1B31]">IELTS</span> exam ?
        </h2>
      </div>
    </div>
  );
}

function slide2() {
  return (
    <div className="flex-1 flex flex-col justify-between items-center text-center">
      {/* Title and Description */}
      <div className="flex-grow flex flex-col justify-center items-center">
        <h2 className="underline decoration-4 underline-offset-15 decoration-red-500 text-4xl font-bold text-white mb-15">
          Welcome to NOAH
        </h2>
        <p className="text-2xl font-bold mt=20 text-white"> Your Personalized IELTS Practice Hub!</p>
      </div>
    </div>
  );
}

function slide3() {
  return (
    <div className="flex-1 flex flex-col justify-between items-center text-center">
      {/* Title and Description */}
      <div className="flex-grow flex flex-col justify-center items-center">
        <h2 className="text-5xl font-bold text-white leading-20">
          Have you booking<br className="mb10"></br> your{' '}
          <span className="text-5xl font-bold text-[#EF1B31]">IELTS</span> Test ?
        </h2>
      </div>
    </div>
  );
}

function slide4() {
  return (
    <div className="flex-1 flex flex-col justify-between items-center text-center">
      {/* Title and Description */}
      <div className="flex-grow flex flex-col justify-center items-center">
        <button className="bg-[#EF1B31] font-semibold text-4xl text-white px-8 py-8 rounded-2xl cursor-pointer">
          Get Start <br></br>Free Trials
        </button>
      </div>
    </div>
  );
}

function SlideBar() {
  const [currentSlide, setCurrentSlide] = useState(0);

  const slides = [
    {
      img: slide_1,
      component: slide1,
    },
    {
      img: slide_2,
      component: slide2,
    },
    {
      img: slide_3,
      component: slide3,
    },
    {
      img: slide_4,
      component: slide4,
    },
  ];

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentSlide((prev) => (prev + 1) % slides.length);
    }, 4000);

    return () => clearInterval(interval);
  }, []);

  const handleSlideChange = (index: number) => {
    setCurrentSlide(index);
  };
  const CurrentComponent = slides[currentSlide].component;

  return (
    <div>
      {/* Slide Bar */}
      <div className="bg-[#070559] px-6 py-16 flex">
        {/* Left Side */}
        <div className="flex-1 flex justify-center items-center">
          {currentSlide === 2 ? (
            <a
              href="https://www.britishcouncil.or.th/en/exam/ielts/dates-fees-locations?utm_source=google&utm_medium=cpc&utm_campaign=exams-all-ielts-ea-th-generic&utm_term=ad2&utm_content=generic-ielts-kw&gad_source=1&gclid=Cj0KCQjwm7q-BhDRARIsACD6-fUUECDaduie7HEUXOJTtH4-3i4KAsKlN348iN_fVGqBOe_Rs3sH44kaAk-mEALw_wcB"
              target="_blank"
              rel="noopener noreferrer"
            >
              <img
                src={slides[currentSlide].img}
                className="lg:w-[500px] h-[350px] object-cover rounded-lg cursor-pointer"
              />
            </a>
          ) : (
            <img src={slides[currentSlide].img} className="lg:w-[500px] h-[350px] object-cover rounded-lg" />
          )}
        </div>

        {/* Right Side */}
        <div className="flex-1 flex flex-col justify-between items-center text-center">
          {/* Title and Description */}
          <div className="flex-grow flex flex-col justify-center items-center">
            <CurrentComponent />
          </div>
          {/* Buttons */}
          <div className="flex justify-center space-x-10 mt-5">
            {slides.map((_, index) => (
              <button
                key={index}
                onClick={() => handleSlideChange(index)}
                className={`px-10 py-1.5 rounded-full cursor-pointer ${
                  currentSlide === index ? 'bg-[#EF1B31]' : 'bg-[#D9D9D9]'
                }`}
              />
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}

function SampleQuestion() {
  return (
    <div>
      <div className="flex justify-center text-4xl font-bold mt-30 mb-30"> View Sample Question</div>
      <div className="flex justify-center space-x-20 mt-20 mb-40">
        <div>
          {' '}
          <img className="w-50 h-50 cursor-pointer" src={listening} />
        </div>
        <div>
          {' '}
          <img className="w-50 h-50 cursor-pointer" src={reading} />
        </div>
        <div>
          {' '}
          <img className="w-50 h-50 cursor-pointer" src={writing} />
        </div>
        <div>
          {' '}
          <img className="w-50 h-50 cursor-pointer" src={speaking} />
        </div>
      </div>
      <div className="flex justify-center w-[1300px] mx-auto h-[7px] bg-[#D9D9D9] mt-20 mb-20 rounded-full"></div>
    </div>
  );
}
const Card = ({ children, className }: { children: React.ReactNode; className?: string }) => {
  return <div className={`rounded-b-sm shadow-2xl bg-white ${className}`}>{children}</div>;
};

const CardContent = ({ children, className }: { children: React.ReactNode; className?: string }) => {
  return <div className={`mt-2 ${className}`}>{children}</div>;
};

function ReviewPage() {
  const [index, setIndex] = useState(0);

  const reviews = [
    [
      {
        img: std_1,
        title: 'grade 11th Student',
        desc: "NOAH ENGLISH makes studying so easy! The personalized practice tests feel like they're tailored just for me. The AI gives detailed feedback that actually helps me improve. Highly recommend it to anyone aiming for a high score!",
      },
      {
        img: std_2,
        title: 'grade 10th Student',
        desc: 'I love how NOAH ENGLISH keeps things interesting! The questions are challenging but not overwhelming, and the mock tests are super realistic. The app feels like a mix of studying and gaming—it makes learning fun instead of stressful.',
      },
      {
        img: std_3,
        title: 'grade 12th Student',
        desc: 'Balancing school and IELTS prep is tough, but NOAH ENGLISH fits into my schedule perfectly. The AI-generated exercises save so much time, and I love how I can choose practice areas based on my weaknesses. It’s like having a personal tutor 24/7!',
      },
    ],
    [
      {
        img: teacher_1,
        title: 'IELTS Senior Grader Review',
        desc: 'NOAH ENGLISH aligns well with the official IELTS format, offering diverse topics and challenges. While the Writing and Speaking modules are solid, more detailed feedback on essay structure and fluency, along with model answers, would enhance its value for teaching.',
      },
      {
        img: teacher_2,
        title: 'IELTS Grader Review',
        desc: 'VThe platform replicates IELTS scoring effectively for grammar and vocabulary but misses some nuances in coherence and task response. Adding detailed feedback on intonation and lexical improvements would make it more accurate.',
      },
      {
        img: teacher_3,
        title: 'IELTS Teacher Review',
        desc: "NOAH ENGLISH does well in simulating test conditions, especially in Reading and Listening. However, Speaking and Writing feedback needs refinement to better match real-world grading. It's a strong tool with room for targeted enhancements.",
      },
    ],
  ];

  useEffect(() => {
    const interval = setInterval(() => {
      setIndex((prevIndex) => (prevIndex + 1) % reviews.length);
    }, 4000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div>
      <div className="w-full max-w-4xl mx-auto text-center mb-20">
        <h2 className="text-3xl font-bold  mb-20">Customer Reviews</h2>
        <div className="overflow-hidden relative">
          <motion.div
            className="flex"
            animate={{ translateX: `-${index * 100}%` }}
            transition={{ duration: 0.8, ease: 'easeInOut' }}
          >
            {reviews.map((frame, i) => (
              <div key={i} className="flex min-w-full gap-4 justify-center">
                {frame.map((review, j) => (
                  <Card key={j} className="p-4 w-1/3">
                    <div className="flex justify-center items-center w-full h-40">
                      <img src={review.img} alt={review.title} className="w-40 h-40 object-cover rounded-md" />
                    </div>
                    <CardContent>
                      <h3 className="font-semibold mt-5 mb-5">{review.title}</h3>
                      <p className="text-sm text-blaack mb-3">{review.desc}</p>
                    </CardContent>
                  </Card>
                ))}
              </div>
            ))}
          </motion.div>
        </div>
      </div>
      <div className="flex justify-center w-[1300px] mx-auto h-[7px] bg-[#D9D9D9] mt-30 mb-20 rounded-full"></div>
    </div>
  );
}

<h2 className="underline decoration-4 underline-offset-15 decoration-red-500 text-4xl font-extrabold text-white mb-15">
  Welcome to NOAH
</h2>;

const Personalization = () => {
  return (
    <div>
      <h1 className=" underline decoration-4 underline-offset-15 decoration-red-500 text-3xl font-bold text-center mb-20">
        Personalization
      </h1>
      <div className="h-130 bg-[#000A52] flex flex-col items-center justify-center text-white rounded-3xl p-6 ml-40 mr-40">
        {/* Header */}

        <div className="flex flex-col md:flex-row w-full max-w-5xl">
          {/* Left Part */}
          <div className="md:w-1/2 flex flex-col justify-center items-center gap-6">
            {['Matching Headings', 'Diagram Label', 'Sentence Completion', 'Short-Answer Question'].map(
              (item, index) => (
                <motion.div
                  key={index}
                  className={`w-64 py-6 text-center rounded-full text-lg font-medium cursor-pointer transition-all relative ${
                    item === 'Diagram Label' ? 'bg-white text-black' : 'bg-white text-black'
                  } ${index % 2 === 0 ? 'self-start' : 'self-end'}`}
                  whileHover={{ scale: 1.05 }}
                  onClick={() => console.log(`${item} clicked`)}
                  style={{ filter: item !== 'Diagram Label' ? 'blur(1.5px)' : 'none' }}
                >
                  {item}
                  {item === 'Diagram Label' && (
                    <div className="absolute -right-6 top-1/2 -translate-y-1/2">
                      <img src={click} className="w-30 h-30 ml-50 mt-15 " />
                    </div>
                  )}
                </motion.div>
              )
            )}
          </div>

          {/* Right Part */}
          <div className="md:w-1/2 flex justify-center items-center text-center ml-auto">
            <p className="text-6xl md:text-4xl text-white pr-10 leading-relaxed ml-35">
              &quot; You can select <br></br>
              <span className="text-red-600 font-bold ">only specific types </span>
              <br></br>
              <span className="text-red-600 font-bold"> of questions</span>
              <span className="text-6xl md:text-4xl  text-white"> you </span>
              <br></br>
              want to focus on &quot;
            </p>
          </div>
        </div>
      </div>
      <div className="flex justify-center w-[1300px] mx-auto h-[7px] bg-[#D9D9D9] mt-30 mb-20 rounded-full"></div>
    </div>
  );
};

function InsightAnalytic() {
  return (
    <div>
      <h1 className="underline decoration-4 underline-offset-15 decoration-red-500 text-3xl font-bold text-center mb-20">
        Insight Analytic
      </h1>
      <div className="flex flex-col md:flex-row justify-center items-center">
        {/* Left Part */}
        <div className="md:w-1/2 flex justify-center items-center text-center md:pl-30">
          <p className="text-6xl md:text-4xl text-black leading-relaxed">
            &quot; Providing a clear <br></br>
            dashboard <span className="text-black font-bold">showing</span> <br></br>
            <span className="text-black font-bold">your strength and</span> <br></br>
            <span className="text-6xl md:text-4xl font-bold text-black"> weaknesses </span> across <br></br>
            all IELTS skills &quot;
          </p>
        </div>

        {/* Right Part */}
        <div className="md:w-1/2 flex justify-center items-center text-center md:pr-30">
          <img src={dashboard} alt="Dashboard" className="w-140 h-140" /> {/* Adjust the size as needed */}
        </div>
      </div>
      <div className="flex justify-center w-[1300px] mx-auto h-[7px] bg-[#D9D9D9] mt-20 mb-20 rounded-full"></div>
    </div>
  );
}

const PricingPage = () => {
  const plans = [
    {
      title: 'Free Explorer',
      price: '0',
      yearlyPrice: '0',
      discount: '',
      features: [
        { text: '4 skills practicing question', available: true },
        { text: '3 practice test', available: true },
        { text: 'Personalization', available: false },
        { text: 'Unlimited Question generated', available: false },
        { text: 'Insight Analytics', available: false },
      ],
    },
    {
      title: 'Pro Learner',
      price: '49',
      beforeprice: '588',
      yearlyPrice: '399',
      discount: '30',
      features: [
        { text: '4 Skills practicing question', available: true },
        { text: '99 practice test/ month', available: true },
        { text: 'Personalization', available: true },
        { text: 'Unlimited Question generated', available: false },
        { text: 'Insight Analytics', available: false },
      ],
    },
    {
      title: 'Premium Mastery',
      price: '79',
      beforeprice: '948',
      yearlyPrice: '599',
      discount: '35',
      features: [
        { text: '4 Skills practicing question', available: true },
        { text: 'Unlimited practice test/ month', available: true },
        { text: 'Personalization', available: true },
        { text: 'Unlimited Question generated', available: true },
        { text: 'Insight Analytics', available: true },
      ],
      isPopular: true,
    },
  ];

  return (
    <div className="min-h-screen bg-[#000A52] py-12 px-4">
      <div className="max-w-7xl mx-auto">
        <div className="text-center mb-12">
          <h1 className="text-4xl font-bold text-white mb-4">Plans that fit your need</h1>
          <p className="text-gray-300">Start with a 7-day free trial.</p>
        </div>

        <div className="grid md:grid-cols-3 gap-8">
          {plans.map((plan, index) => (
            <div
              key={index}
              className={`bg-white p-6 rounded-lg shadow-lg ${
                plan.isPopular ? 'border-2 border-red-500 relative' : ''
              }`}
            >
              {plan.isPopular && (
                <div className="absolute top-0 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
                  <span className="bg-red-500 text-white px-5 py-2 rounded-full text-sm">Most popular</span>
                </div>
              )}
              <h2 className="text-2xl font-bold text-navy-900 mb-15 underline decoration-4 underline-offset-10 decoration-red-500">
                {plan.title}
              </h2>
              <div className="mb-15 flex flex-col items-center">
                <p className="text-xl">
                  <span className="text-2xl">{plan.price}</span> Bath/month
                </p>
                <p className="text-xl">
                  <del>{plan.beforeprice}</del> <span className="text-4xl font-bold">{plan.yearlyPrice}</span> Bath/year
                </p>
                {plan.discount && <p className="text-2xl font-bold text-red-500">Save {plan.discount} % </p>}
              </div>

              <div className="space-y-3">
                {plan.features.map((feature, featureIndex) => (
                  <div key={featureIndex} className="flex items-center">
                    {feature.available ? (
                      <svg className="w-5 h-5 text-green-500 mr-2" fill="currentColor" viewBox="0 0 20 20">
                        <path
                          fillRule="evenodd"
                          d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                          clipRule="evenodd"
                        />
                      </svg>
                    ) : (
                      <svg className="w-5 h-5 text-red-500 mr-2" fill="currentColor" viewBox="0 0 20 20">
                        <path
                          fillRule="evenodd"
                          d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                          clipRule="evenodd"
                        />
                      </svg>
                    )}
                    <span className="text-gray-700">{feature.text}</span>
                  </div>
                ))}
              </div>
              <div className="flex justify-center">
                <button
                  className="mt-6 w-1/2 bg-white text-[#000A52] text-center py-2 px-4 rounded border border-[#000A52] hover:bg-[#000A52] hover:text-white transition-colors"
                  onClick={() => console.log(`Selected ${plan.title}`)}
                >
                  Get Started
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};
