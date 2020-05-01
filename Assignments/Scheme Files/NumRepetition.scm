#lang scheme
;check if input is list
(define (sameNum mylist)
         (if (list? mylist)
             (check-repeated mylist)
             'list-error))

(define (check-repeated list2)
  (let ((c1 1) (greatCount 0) (greatVar 0))
    (let recursive ((mylist list2))
          (cond
           	((null? mylist) '())
                ((null? (cdr mylist)) (cond 
                ((>= c1 greatCount) (set! greatVar (car mylist))  (set! greatCount c1) '())))           
                ((= (car mylist) (cadr mylist)) (set! c1(add1 c1)) (recursive (cdr mylist)))
						((not (= (car mylist) (cadr mylist))) 
             (cond 
             ((>= c1 greatCount) (set! greatCount c1) (set! greatVar (car mylist))  (set! c1 1) (recursive (cdr mylist)))
             (else (set! c1 1) (recursive (cdr mylist)))))))
         (let ((newlist '()))
         (for/list ((i greatCount))  (set! newlist(append (list greatVar) newlist))) newlist)
       ))
     
              
              
						
          

  
          
            

        


     
 



           
