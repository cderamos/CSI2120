#lang scheme
(require math/base)

;a
(define sigmoid (lambda (value) (/ 1 (+ 1 (/ 1 (expt euler.0 value))))))

(define (zetaCalc zetaParams inputList)
        (+ (car zetaParams) (* (cadr zetaParams) (car inputList)) (* (car (reverse zetaParams)) (cadr inputList))))

(define (thetaCalc thetaParams zetaList)
        (+ (car thetaParams) (* (car zetaList) (cadr thetaParams))
           (*  (cadr zetaList) (car (cdr (cdr thetaParams)))) (* (car (cdr (cdr zetaList))) (car (reverse thetaParams)))))

(define (neuralNode myList someFunc)
        (lambda (inputVals)
          (cond 
          ((= (length myList) 3) (someFunc (zetaCalc myList inputVals)))
          ((= (length myList) 4) (someFunc (thetaCalc myList inputVals))))))

;b
(define (neuralLayer moreZetaParams)
        (lambda (inputVals) (list (sigmoid (zetaCalc (car moreZetaParams) inputVals))
                                  (sigmoid (zetaCalc (cadr moreZetaParams) inputVals))
                                  (sigmoid (zetaCalc (car (reverse moreZetaParams)) inputVals)))))

;c
(define (neuralNet inputs)
  ((neuralNode '(0.5 0.3 0.7 0.1) sigmoid) ((neuralLayer '((0.1 0.3 0.4)(0.5 0.8 0.3)(0.7 0.6 0.6))) inputs) ))

;d
(define (applyNet itera)
  (reverse ((lambda (itera) (let recur ((var itera))
    (if (= var 0) '()
        (cons (neuralNet (list (sin (/ (* 2 pi (- var 1)) itera)) (cos (/ (* 2 pi (- var 1)) itera)))) (recur (- var 1)))))) itera )))

