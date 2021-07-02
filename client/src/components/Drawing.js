import { useEffect, useRef } from 'react';
import './Drawing.css';


const Drawing = ({pen, canvas}) => {
    const canvasRef = useRef(null);
    const ctxRef = useRef(null);
    const animationRef = useRef(null);
    const canvasModified = useRef(false);
    const mousePosition = useRef(null);
    const mouseDown = useRef(false);
    const pathStarted = useRef(false);

    const getCanvasCoordsFromEvent = (e) => {
        if (canvasRef.current === null || e.target !== canvasRef.current)
            return null;

        return {
            x: e.pageX - canvasRef.current.offsetLeft,
            y: e.pageY - canvasRef.current.offsetTop,
        }
    }

    const updateDrawing = () => {
        if (ctxRef.current === null || mousePosition.current === null || mouseDown.current === false) {
            animationRef.current = requestAnimationFrame(updateDrawing);
            return;
        }
        
        const ctx = ctxRef.current;
        const pos = mousePosition.current;
        if (pathStarted.current === false) {
            ctx.beginPath();
            ctx.moveTo(pos.x, pos.y);
            pathStarted.current = true;
        } else {
            ctx.lineTo(pos.x, pos.y);
            ctx.stroke();
        }
        canvasModified.current = true;
        animationRef.current = requestAnimationFrame(updateDrawing);
    }

    useEffect(() => {
        if (canvasRef.current === null) {
            ctxRef.current = null;
            canvas.current = null;
            return;
        }
        
        ctxRef.current = canvasRef.current.getContext('2d');
        ctxRef.current.lineCap = 'round';
        ctxRef.current.lineJoin = 'round';
        canvas.current = ctxRef.current;
    }, [canvasRef, canvas]);

    useEffect(() => {
        if (ctxRef.current === null)
            return;

        ctxRef.current.strokeStyle = pen.color;
        ctxRef.current.lineWidth = pen.thickness;
    }, [ctxRef, pen]);

    const startTrackingMouse = (e) => {
        mousePosition.current = getCanvasCoordsFromEvent(e);
        animationRef.current = requestAnimationFrame(updateDrawing);
    }

    const stopTrackingMouse = () => {
        cancelAnimationFrame(animationRef.current);
        pathStarted.current = false;
        mousePosition.current = null;
    }

    const handleMouseDown = (e) => {
        mouseDown.current = true;
        pathStarted.current = false;
        mousePosition.current = getCanvasCoordsFromEvent(e);
    }

    const handleMouseMove = (e) => {
        if (mouseDown.current === false)
            return;
        
        mousePosition.current = getCanvasCoordsFromEvent(e);
    }

    useEffect(() => {
        const handleMouseUp = () => {
            mouseDown.current = false;
            if (canvasModified.current === true && ctxRef.current !== null) {
                ctxRef.current.save();
            }
            canvasModified.current = false;
        }
        document.addEventListener('mouseup', handleMouseUp);

        return () => {
            document.removeEventListener('mouseup', handleMouseUp);
        }
    }, [mouseDown, canvasModified, ctxRef]);

    return (
        <canvas
            ref={canvasRef}
            onMouseEnter={startTrackingMouse}
            onMouseLeave={stopTrackingMouse}
            onMouseDown={handleMouseDown}
            onMouseMove={handleMouseMove}
            width={400}
            height={400}
            className="Drawing"></canvas>
    )
}

export default Drawing;
