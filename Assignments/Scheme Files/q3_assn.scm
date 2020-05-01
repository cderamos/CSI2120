#lang scheme
(define choices '(("marie" ("peru" "greece" "vietnam"))
                  ("jean" ("greece" "peru" "vietnam"))
                  ("sasha" ("vietnam" "peru" "greece"))
                  ("helena" ("peru" "vietnam" "greece"))
                  ("emma" ("greece" "peru" "vietnam"))))

;check if input is list
(define (chooseTrip mylist)
         (if (list? mylist)
             (calculateTrip mylist)
             'list-error))

(define (calculateTrip mylist)
         (let ((countries (list-ref (list-ref mylist 0) 1)) (count 1));get a list of all the countries (peru greece vietnam)
               (define ht (make-hash)) ;making a hash table from that list of countries
               (let ( (c1 (car countries)) (c2 (cadr countries)) (c3 (car (reverse countries)))) ;set to vars to make it easier to read
               (hash-set! ht c1 0) (hash-set! ht c2 0) (hash-set! ht c3 0)
               (let recur ((list mylist))
                   (cond
                       ((not(null? list))
                        (let recur2 ((list2 (car (cdr (car list)))) (points 3))
                             (cond
                                  ((not (null? list2)) (cond
                                  ((eq?  (car list2) c1) (hash-set! ht c1 (+ points (hash-ref ht c1))) (recur2 (cdr list2) (- points 1)))
                                  ((eq?  (car list2) c2) (hash-set! ht c2 (+ points (hash-ref ht c2))) (recur2 (cdr list2) (- points 1)))
                                  ((eq?  (car list2) c3) (hash-set! ht c3 (+ points (hash-ref ht c3))) (recur2 (cdr list2) (- points 1)))))))
                        (recur (cdr list))))))
                        (cond
                           ((=(car(hash-values ht)) (cadr (hash-values ht)) (car (cdr (cdr (hash-values ht))))) ht)
                           ((=(car(hash-values ht)) (cadr (hash-values ht))) (list (car(hash->list ht)) (cadr (hash-values ht)))) 
                           ((=(car(hash-values ht)) (car (cdr (cdr (hash-values ht)))))  (list (car (hash->list ht)) (car(cdr(cdr (hash->list ht))))) )
                           ((=(cadr (hash-values ht)) (car (cdr (cdr (hash-values ht))))) (cdr(hash->list ht)))
                           (else
                              (argmax cdr (hash->list ht))))))        
           