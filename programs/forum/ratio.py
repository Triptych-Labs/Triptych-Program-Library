if __name__ == "__main__":
    print("....")

    per = 40
    exch = 1
    C = 80

    ### Compute D of C where A:B
    # divide A by B
    BdivsA = exch / per
    # multiply C by (B/A)
    D = C * BdivsA

    print(D)

